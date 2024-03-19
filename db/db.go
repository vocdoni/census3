package db

import (
	"bufio"
	"bytes"
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	queries "github.com/vocdoni/census3/db/sqlc"
	"go.vocdoni.io/dvote/log"
)

//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.23.0 generate

//go:embed migrations/*.sql
var migrationsFS embed.FS

// DB struct abstact a safe connection with the database using sqlc queries,
// sqlite as a database engine and go-sqlite3 as a driver.
type DB struct {
	RW *sql.DB
	RO *sql.DB

	QueriesRW *queries.Queries
	QueriesRO *queries.Queries
}

// Close function stops all internal connections to the database
func (db *DB) Close() error {
	if err := db.RW.Close(); err != nil {
		return err
	}
	return db.RO.Close()
}

// Init function starts a database using the data path provided as argument. It
// opens two different connections, one for read only, and another for read and
// write, with different configurations, optimized for each use case.
func Init(dataDir string, dbName string) (*DB, error) {
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

func (db *DB) Import(ctx context.Context, dump []byte) error {
	tx, err := db.RW.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Errorw(err, "error rolling back transaction")
		}
	}()
	reader := bytes.NewReader(dump)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		sql := scanner.Text()
		if sql == "" {
			continue
		}
		sql = strings.TrimSpace(sql)
		log.Info(sql)
		if _, err := tx.ExecContext(ctx, sql); err != nil {
			log.Warn(sql)
			return fmt.Errorf("error executing SQL: %w", err)
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}
	return nil
}

func (db *DB) Export(ctx context.Context) ([]byte, error) {
	dumpSQL := []byte{}
	buf := bytes.NewBuffer(dumpSQL)
	tokens, err := db.exportTokens(ctx)
	if err != nil {
		return nil, fmt.Errorf("error exporting tokens: %w", err)
	}
	buf.Write(tokens)
	holders, err := db.exportHolders(ctx)
	if err != nil {
		return nil, fmt.Errorf("error exporting holders: %w", err)
	}
	buf.Write(holders)
	return buf.Bytes(), nil
}

func (db *DB) exportTokens(ctx context.Context) ([]byte, error) {
	tokens, err := db.QueriesRO.DumpTokens(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting tokens: %w", err)
	}
	resultSQL := []byte{}
	buf := bytes.NewBuffer(resultSQL)
	for _, token := range tokens {
		// get default strategy
		defaultStrategy, err := db.QueriesRO.StrategyByID(ctx, token.DefaultStrategy)
		if err != nil {
			return nil, fmt.Errorf("error getting default strategy: %w", err)
		}
		strategyTokens, err := db.QueriesRO.StrategyTokens(ctx, token.DefaultStrategy)
		if err != nil {
			return nil, fmt.Errorf("error getting strategy tokens: %w", err)
		}
		bSynced := 0
		if token.Synced {
			bSynced = 1
		}
		tokenInsert := fmt.Sprintf("INSERT INTO tokens "+
			"(id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, "+
			"chain_id, chain_address, external_id, default_strategy, icon_uri, last_block) "+
			"VALUES (X'%x', '%s', '%s', %d, '%s', %d, %d, %d, '%s', %d, '%s', '%s', %d, '%s', %d);\n",
			token.ID, token.Name, token.Symbol, token.Decimals, token.TotalSupply, token.CreationBlock,
			token.TypeID, bSynced, token.Tags, token.ChainID, token.ChainAddress, token.ExternalID,
			token.DefaultStrategy, token.IconUri, token.LastBlock)
		strategyInsert := fmt.Sprintf("INSERT INTO strategies "+
			"(id, predicate, alias, uri) VALUES (%d, '%s', '%s', '%s');\n",
			defaultStrategy.ID, defaultStrategy.Predicate, defaultStrategy.Alias, defaultStrategy.Uri)
		strategyTokensInsert := ""
		for _, st := range strategyTokens {
			strategyTokensInsert += fmt.Sprintf("INSERT INTO strategy_tokens "+
				"(strategy_id, token_id, min_balance, chain_id, external_id) "+
				"VALUES (%d, X'%x', '%s', %d, '%s');\n",
				defaultStrategy.ID, st.TokenID, st.MinBalance, st.ChainID, st.ExternalID)
		}
		// write to buffer
		buf.WriteString(tokenInsert)
		buf.WriteString(strategyInsert)
		buf.WriteString(strategyTokensInsert)
	}
	return buf.Bytes(), nil
}

func (db *DB) exportHolders(ctx context.Context) ([]byte, error) {
	holders, err := db.QueriesRW.DumpTokenHolers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting tokens: %w", err)
	}
	resultSQL := []byte{}
	buf := bytes.NewBuffer(resultSQL)
	for _, holder := range holders {
		insert := fmt.Sprintf("INSERT INTO token_holders "+
			"(token_id, holder_id, balance, block_id, chain_id, external_id) "+
			"VALUES (X'%x', X'%x', '%s', %d, %d, '%s');\n",
			holder.TokenID, holder.HolderID, holder.Balance, holder.BlockID, holder.ChainID, holder.ExternalID)
		buf.WriteString(insert)
	}
	return buf.Bytes(), nil
}
