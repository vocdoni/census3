package db

import (
	"database/sql"
	"embed"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	queries "github.com/vocdoni/census3/db/sqlc"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func Init(dataDir string) (*sql.DB, *queries.Queries, error) {
	dbFile := filepath.Join(dataDir, "census3.sql")
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
			return nil, nil, fmt.Errorf("error creating a new database file: %w", err)
		}
	}
	// open database file
	database, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening database: %w", err)
	}
	// get census3 goose migrations and setup for sqlite3
	goose.SetDialect("sqlite3")
	goose.SetBaseFS(migrationsFS)
	// perform goose up
	if err := goose.Up(database, "migrations"); err != nil {
		return nil, nil, fmt.Errorf("error during goose up: %w", err)
	}
	// init sqlc
	return database, queries.New(database), nil
}
