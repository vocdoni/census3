package farcaster

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pressly/goose/v3"
	queries "github.com/vocdoni/census3/scanner/providers/farcaster/sqlc"
	"go.vocdoni.io/dvote/log"
)

// DB struct abstact a safe connection with the database using sqlc queries,
// sqlite as a database engine and go-sqlite3 as a driver.
type DB struct {
	RW *sql.DB
	RO *sql.DB

	QueriesRW *queries.Queries
	QueriesRO *queries.Queries
}

// Close function stops all internal connections to the database
func (db *DB) CloseDB() error {
	if err := db.RW.Close(); err != nil {
		return err
	}
	return db.RO.Close()
}

// Init function starts a database using the data path provided as argument. It
// opens two different connections, one for read only, and another for read and
// write, with different configurations, optimized for each use case.
func InitDB(dataDir string, dbName string) (*DB, error) {
	if dbName == "" {
		return nil, fmt.Errorf("database name is required")
	}
	dbFile := filepath.Join(dataDir, dbName)
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
			return nil, fmt.Errorf("error creating a new database file: %w", err)
		}
	}
	// sqlite doesn't support multiple concurrent writers.
	// For that reason, rwDB is limited to one open connection.
	// Per https://github.com/mattn/go-sqlite3/issues/1022#issuecomment-1067353980,
	// we use WAL to allow multiple concurrent readers at the same time.
	rwDB, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=rwc&_journal_mode=wal&_txlock=immediate&_synchronous=normal", dbFile))
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	rwDB.SetMaxOpenConns(1)
	rwDB.SetMaxIdleConns(2)
	rwDB.SetConnMaxIdleTime(10 * time.Minute)
	rwDB.SetConnMaxLifetime(time.Hour)

	roDB, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=ro&_journal_mode=wal", dbFile))
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	// Increasing these numbers can allow for more queries to run concurrently,
	// but it also increases the memory used by sqlite and our connection pool.
	// Most read-only queries we run are quick enough, so a small number seems OK.
	roDB.SetMaxOpenConns(10)
	roDB.SetMaxIdleConns(20)
	roDB.SetConnMaxIdleTime(5 * time.Minute)
	roDB.SetConnMaxLifetime(time.Hour)

	// get census3 goose migrations and setup for sqlite3
	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, fmt.Errorf("error setting up driver for sqlite: %w", err)
	}
	goose.SetBaseFS(migrationsFS)
	// perform goose up
	if err := goose.Up(rwDB, "migrations"); err != nil {
		return nil, fmt.Errorf("error during goose up: %w", err)
	}
	// init sqlc
	return &DB{
		RW:        rwDB,
		RO:        roDB,
		QueriesRW: queries.New(rwDB),
		QueriesRO: queries.New(roDB),
	}, nil
}

// updates farcaster database with the users data
func (p *FarcasterProvider) updateFarcasterDB(ctx context.Context, usersData []*FarcasterUserData) error {
	// init db transaction
	internalCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	tx, err := p.db.RW.BeginTx(internalCtx, nil)
	if err != nil {
		return fmt.Errorf("cannot update farcaster db: %w", err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Errorw(err, "farcaster transaction rollback failed")
		}
	}()

	qtx := p.db.QueriesRW.WithTx(tx)
	// iterate the users and update the database
	// if the user does not exist on the database create a new one
	// if it exists all info except appkeys is already stored, so just update appkeys
	for _, userData := range usersData {
		// check if the user exists
		user, errGetUser := qtx.GetUserByFID(internalCtx, userData.FID.Uint64())
		if errGetUser != nil {
			// if not exists create a new user
			if errors.Is(errGetUser, sql.ErrNoRows) {
				if err := p.createUser(internalCtx, qtx, userData); err != nil {
					return fmt.Errorf("cannot create user %w", err)
				}
			} else {
				return fmt.Errorf("cannot update farcaster db: %w", errGetUser)
			}
		} else { // if user exists update the user app keys, all other data is already stored
			if err := p.updateUserAppKeys(internalCtx, qtx, user, userData.AppKeys); err != nil {
				return fmt.Errorf("cannot update user data %w", err)
			}
		}
		if err := p.createLinkedEVMFID(ctx, qtx, userData.LinkedEVM, userData.FID.Uint64()); err != nil {
			return fmt.Errorf("cannot update farcaster db: %w", err)
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("cannot update farcaster db: %w", err)
	}
	return nil
}

func (p *FarcasterProvider) createUser(ctx context.Context, qtx *queries.Queries, userData *FarcasterUserData) error {
	serializedLinkedEVM := make([]byte, 0)
	var err error
	if len(userData.LinkedEVM) != 0 {
		serializedLinkedEVM, err = serializeArray(userData.LinkedEVM)
		if err != nil {
			return fmt.Errorf("cannot update farcaster db: %w", err)
		}
	}
	if _, err := qtx.CreateUser(ctx, queries.CreateUserParams{
		Fid: userData.FID.Uint64(),
		// Username:        userData.Username,
		Signer:          userData.Signer[:],
		CustodyAddress:  userData.CustodyAddress[:],
		RecoveryAddress: userData.RecoveryAddress[:],
		LinkedEvm:       serializedLinkedEVM,
		AppKeys:         make([]byte, 0),
	}); err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fmt.Errorf("cannot update farcaster db: %w", ErrUserAlreadyExists)
		}
		return fmt.Errorf("cannot update farcaster db: %w", err)
	}
	return nil
}

func (p *FarcasterProvider) createLinkedEVMFID(
	ctx context.Context, qtx *queries.Queries, linkedEVM []common.Address, fid uint64,
) error {
	for _, evmKey := range linkedEVM {
		if _, err := qtx.CreateLinkedEVMFID(ctx, queries.CreateLinkedEVMFIDParams{
			Fid:        fid,
			EvmAddress: evmKey[:],
		}); err != nil {
			return fmt.Errorf("cannot update farcaster db: %w", err)
		}
	}
	return nil
}

func (p *FarcasterProvider) updateUserAppKeys(
	ctx context.Context, qtx *queries.Queries, user queries.User, appKeys []common.Hash,
) error {
	// serialize app keys before saving
	serializedAppKeys := make([]byte, 0)
	var err error
	if len(appKeys) != 0 {
		serializedAppKeys, err = serializeArray(appKeys)
		if err != nil {
			return fmt.Errorf("cannot update farcaster db: %w", err)
		}
	}
	if _, err := qtx.UpdateUserAppKeys(ctx, queries.UpdateUserAppKeysParams{
		Fid:     user.Fid,
		AppKeys: serializedAppKeys,
	}); err != nil {
		return fmt.Errorf("cannot update farcaster db: %w", err)
	}
	return nil
}
