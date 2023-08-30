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
	fileURI := fmt.Sprintf("file:%s?cache=shared", dbFile)
	database, err := sql.Open("sqlite3", fileURI)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening database: %w", err)
	}
	// trying to fix "database is locked" issue according to the official
	// mattn/go-sqlite3 docs: https://github.com/mattn/go-sqlite3/#faq
	database.SetMaxOpenConns(1)
	// get census3 goose migrations and setup for sqlite3
	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, nil, fmt.Errorf("error setting up driver for sqlite: %w", err)
	}
	goose.SetBaseFS(migrationsFS)
	// perform goose up
	if err := goose.Up(database, "migrations"); err != nil {
		return nil, nil, fmt.Errorf("error during goose up: %w", err)
	}
	// init sqlc
	return database, queries.New(database), nil
}
