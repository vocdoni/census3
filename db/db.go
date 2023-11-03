package db

import (
	"database/sql"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	queries "github.com/vocdoni/census3/db/sqlc"
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
func Init(dataDir string) (*DB, error) {
	dbFile := filepath.Join(dataDir, "census3.sql")
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
