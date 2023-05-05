package db

import (
	"database/sql"
	"embed"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	postgresMigrations "github.com/vocdoni/census3/db/postgres"
	queries "github.com/vocdoni/census3/db/sqlc"
	sqlite3Migrations "github.com/vocdoni/census3/db/sqlite3"
	"go.vocdoni.io/dvote/log"
)

type DatabaseEngine int32

const (
	PostgresEngine DatabaseEngine = iota
	SQLiteEngine
)

var EngineToString = map[DatabaseEngine]string{
	PostgresEngine: "postgres",
	SQLiteEngine:   "sqlite3",
}

var StringToEngine = map[string]DatabaseEngine{
	"postgres": PostgresEngine,
	"sqlite3":  SQLiteEngine,
}

func Init(engine DatabaseEngine, dataDir string) (*queries.Queries, error) {
	var model embed.FS
	var connStr string
	switch engine {
	case PostgresEngine:
		connStr = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_DB"))
		model = postgresMigrations.Migrations
	case SQLiteEngine:
		connStr = filepath.Join(dataDir, "census3.sql")
		if _, err := os.Stat(connStr); os.IsNotExist(err) {
			if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
				return nil, fmt.Errorf("error creating a new database file: %w", err)
			}
		}
		model = sqlite3Migrations.Migrations
	default:
		return nil, fmt.Errorf("database engine not implemented")
	}
	log.Infow("database starting...", "engine", EngineToString[engine], "conn", connStr)
	// connect to the database
	database, err := sql.Open(EngineToString[engine], connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	// set census3 goose migrations
	goose.SetBaseFS(model)
	// setup goose for the engine
	if err := goose.SetDialect(EngineToString[engine]); err != nil {
		return nil, fmt.Errorf("error during goose dialect setup: %w", err)
	}
	// perform goose up
	if err := goose.Up(database, "migrations"); err != nil {
		return nil, fmt.Errorf("error during goose up: %w", err)
	}
	// init sqlc
	return queries.New(database), nil
}
