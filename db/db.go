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
	queries "github.com/vocdoni/census3/db/sqlc"
	"go.vocdoni.io/dvote/log"
)

//go:embed migrations/sqlite3/*.sql
var migrationsSQLite3 embed.FS

//go:embed migrations/postgres/*.sql
var migrationsPostgres embed.FS

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
	var migrations embed.FS
	var connStr string
	switch engine {
	case PostgresEngine:
		connStr = fmt.Sprintf("postgres://%s@%s/%s?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_DB"))
		// connStr = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		// 	os.Getenv("POSTGRES_USER"),
		// 	os.Getenv("POSTGRES_PASSWORD"),
		// 	os.Getenv("POSTGRES_HOST"),
		// 	os.Getenv("POSTGRES_DB"))
		migrations = migrationsPostgres
	case SQLiteEngine:
		connStr = filepath.Join(dataDir, "census3.sql")
		if _, err := os.Stat(connStr); os.IsNotExist(err) {
			if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
				return nil, fmt.Errorf("error creating a new database file: %w", err)
			}
		}
		migrations = migrationsSQLite3
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
	goose.SetBaseFS(migrations)
	// setup goose for the engine
	goose.SetDialect(EngineToString[engine])
	// perform goose up
	if err := goose.Up(database, "migrations"); err != nil {
		return nil, fmt.Errorf("error during goose up: %w", err)
	}
	// init sqlc
	return queries.New(database), nil
}
