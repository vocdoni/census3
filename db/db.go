package db

import (
	"bytes"
	"context"
	"database/sql"
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
	path string

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
	// init sqlc
	db := &DB{path: filepath.Join(dataDir, dbName)}
	if err := db.createIfNotExists(dataDir); err != nil {
		return nil, fmt.Errorf("error creating database: %w", err)
	}
	if err := db.start(); err != nil {
		return nil, fmt.Errorf("error starting database: %w", err)
	}
	return db, nil
}

func (db *DB) createIfNotExists(dir string) error {
	if _, err := os.Stat(db.path); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("error creating a new database file: %w", err)
		}
	}
	return nil
}

func (db *DB) start() error {
	var err error
	// sqlite doesn't support multiple concurrent writers.
	// For that reason, rwDB is limited to one open connection.
	// Per https://github.com/mattn/go-sqlite3/issues/1022#issuecomment-1067353980,
	// we use WAL to allow multiple concurrent readers at the same time.
	db.RW, err = sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=rwc&_journal_mode=wal&_txlock=immediate&_synchronous=normal", db.path))
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}
	db.RW.SetMaxOpenConns(1)
	db.RW.SetMaxIdleConns(2)
	db.RW.SetConnMaxIdleTime(10 * time.Minute)
	db.RW.SetConnMaxLifetime(time.Hour)

	db.RO, err = sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=ro&_journal_mode=wal", db.path))
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}
	// Increasing these numbers can allow for more queries to run concurrently,
	// but it also increases the memory used by sqlite and our connection pool.
	// Most read-only queries we run are quick enough, so a small number seems OK.
	db.RO.SetMaxOpenConns(10)
	db.RO.SetMaxIdleConns(20)
	db.RO.SetConnMaxIdleTime(5 * time.Minute)
	db.RO.SetConnMaxLifetime(time.Hour)
	// get census3 goose migrations and setup for sqlite3
	if err := goose.SetDialect("sqlite3"); err != nil {
		return fmt.Errorf("error setting up driver for sqlite: %w", err)
	}
	goose.SetBaseFS(migrationsFS)
	// perform goose up
	if err := goose.Up(db.RW, "migrations"); err != nil {
		return fmt.Errorf("error during goose up: %w", err)
	}
	db.QueriesRW = queries.New(db.RW)
	db.QueriesRO = queries.New(db.RO)
	return nil
}

// Import function replaces the current database with the one provided as a
// byte array. It creates a temporary file to store the dump, and then copies
// the dump to the database file. Finally, it starts the database again. If any
// error occurs, it returns an error.
func (db *DB) Import(ctx context.Context, dump []byte) error {
	target, err := os.Create(db.path)
	if err != nil {
		return fmt.Errorf("error creating database file: %w", err)
	}
	if _, err := io.Copy(target, bytes.NewReader(dump)); err != nil {
		return fmt.Errorf("error copying dump to database file: %w", err)
	}
	if err := target.Close(); err != nil {
		return fmt.Errorf("error closing database file: %w", err)
	}
	return db.start()
}

// Export function returns the current database as a byte array. It creates a
// temporary file to store the dump, and then copies the database to the dump
// file. Finally, it reads the dump file and returns it as a byte array. If any
// error occurs, it returns an error. It uses the VACUUM INTO command to create
// a dump of the database.
func (db *DB) Export(ctx context.Context) ([]byte, error) {
	tempDir, err := os.MkdirTemp("", "census3-export")
	if err != nil {
		return nil, fmt.Errorf("error creating temp directory: %w", err)
	}
	dbDump := filepath.Join(tempDir, "census3.sqlite3")
	if _, err := db.RO.ExecContext(ctx, "VACUUM INTO ?", dbDump); err != nil {
		return nil, fmt.Errorf("error vacuuming database: %w", err)
	}
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			log.Errorw(err, "error removing temp directory")
		}
	}()
	return os.ReadFile(dbDump)
}
