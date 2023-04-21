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

func Init(dataDir string) (*queries.Queries, error) {
	dbFile := filepath.Join(dataDir, "census3.sql")
	if _, err := os.Stat(dbFile); err != nil {
		fd, err := os.Create(dbFile)
		if err != nil {
			return nil, fmt.Errorf("error creating a new database file: %w", err)
		}
		fd.Close()
	}

	// open database file
	database, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	fmt.Println(filepath.Join(dataDir, "census3.sql"))
	// get census3 goose migrations and setup for sqlite3
	goose.SetDialect("sqlite3")
	goose.SetBaseFS(migrationsFS)
	// perform goose up
	if err := goose.Up(database, "migrations"); err != nil {
		return nil, fmt.Errorf("error during goose up: %w", err)
	}
	// init sqlc
	return queries.New(database), nil
}
