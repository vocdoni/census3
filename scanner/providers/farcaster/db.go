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
func (p *FarcasterProvider) updateFarcasterDB(ctx context.Context, usersData []FarcasterUserData) error {
	// init db transaction
	internalCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	tx, err := p.db.RW.BeginTx(internalCtx, nil)
	if err != nil {
		return fmt.Errorf("cannot update farcaster db: %w", err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Errorw(err, "farcaster transaction rollback failed")
		}
	}()

	qtx := p.db.QueriesRW.WithTx(tx)
	// iterate the users and update the database
	// if the user does not exist on the database create a new one
	// if it exists all info except appkeys is already stored, so just update appkeys
	for _, userData := range usersData {
		// check if the user exists
		_, err := qtx.GetUserByFID(internalCtx, userData.FID)
		if err != nil {
			// if not exists create a new user
			if errors.Is(err, sql.ErrNoRows) {
				if err := p.createUser(internalCtx, qtx, userData); err != nil {
					return fmt.Errorf("cannot create user %w", err)
				}
			} else {
				return fmt.Errorf("cannot update farcaster db: %w", err)
			}
		}
		// if err := p.createLinkedEVMFID(ctx, qtx, userData.LinkedEVM, userData.FID.Uint64()); err != nil {
		// 	return fmt.Errorf("cannot update farcaster db: %w", err)
		// }
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("cannot update farcaster db: %w", err)
	}
	return nil
}

func (p *FarcasterProvider) createUser(ctx context.Context, qtx *queries.Queries, userData FarcasterUserData) error {
	if _, err := qtx.CreateUser(ctx, userData.FID); err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fmt.Errorf("cannot update farcaster db: %w", ErrUserAlreadyExists)
		}
		return fmt.Errorf("cannot update farcaster db: %w", err)
	}
	return nil
}

func (p *FarcasterProvider) createFidAppKey(
	ctx context.Context, fid uint64, appKey common.Hash,
) error {
	if _, err := p.db.QueriesRW.CreateFidAppKey(ctx, queries.CreateFidAppKeyParams{
		Fid:    fid,
		AppKey: appKey[:],
	}); err != nil {
		return fmt.Errorf("cannot update farcaster db: %w", err)
	}
	return nil
}

func (p *FarcasterProvider) deleteFidAppKey(
	ctx context.Context, fid uint64, appKey common.Hash,
) error {
	if _, err := p.db.QueriesRW.DeleteFidAppKey(ctx, queries.DeleteFidAppKeyParams{
		Fid:    fid,
		AppKey: appKey[:],
	}); err != nil {
		return fmt.Errorf("cannot update farcaster db: %w", err)
	}
	return nil
}

func (p *FarcasterProvider) addAppKeys(
	ctx context.Context, fidList []uint64, addedKeys map[uint64][][]byte,
) error {
	if len(fidList) == 0 {
		return nil
	}
	for _, fid := range fidList {
		keys, ok := addedKeys[fid]
		if !ok {
			continue
		}
		for _, key := range keys {
			exists, err := p.db.QueriesRO.CheckFidAppKeyExists(ctx, queries.CheckFidAppKeyExistsParams{
				Fid:    fid,
				AppKey: key[:],
			})
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("cannot get fid app keys %w", err)
			} else if exists {
				continue
			}
			h := common.Hash{}
			h.SetBytes(key)
			// create ref for each fid and key on fid_appkeys table
			if err := p.createFidAppKey(ctx, fid, h); err != nil {
				return fmt.Errorf("cannot create fid app key %w", err)
			}
		}
	}
	return nil
}

func (p *FarcasterProvider) deleteAppKeys(
	ctx context.Context, fidList []uint64, deletedKeys map[uint64][][]byte,
) error {
	if len(fidList) == 0 {
		return nil
	}
	for _, fid := range fidList {
		keys, ok := deletedKeys[fid]
		if !ok {
			continue
		}
		for _, key := range keys {
			exists, err := p.db.QueriesRO.CheckFidAppKeyExists(ctx, queries.CheckFidAppKeyExistsParams{
				Fid:    fid,
				AppKey: key[:],
			})
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("cannot get fid app keys %w", err)
			} else if exists {
				h := common.Hash{}
				h.SetBytes(key)
				if err := p.deleteFidAppKey(ctx, fid, h); err != nil {
					return fmt.Errorf("cannot delete fid app key %w", err)
				}
			}
		}
	}
	return nil
}
