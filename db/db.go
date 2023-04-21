package db

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"path"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	queries "github.com/vocdoni/census3/db/sqlc"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func Init(dataDir string) (*queries.Queries, error) {
	log.SetFlags(0)
	// open database file
	database, err := sql.Open("sqlite3", filepath.Join(dataDir, "census3.sql"))
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	// get census3 goose migrations and setup for sqlite3
	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, fmt.Errorf("error during setup goose: %w", err)
	}
	// perform goose up
	goose.SetBaseFS(migrationsFS)

	sqlMigrationFiles, err := fs.Glob(migrationsFS, path.Join("migrations", "*.sql"))
	if err != nil {
		return nil, err
	}
	for _, i := range sqlMigrationFiles {
		fmt.Println(i)
	}

	if err := goose.Up(database, "migrations"); err != nil {
		return nil, fmt.Errorf("error during goose up: %w", err)
	}
	// init sqlc
	return queries.New(database), nil
}
