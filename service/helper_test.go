package service

import (
	"database/sql"
	"os"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/census3/db"
	queries "github.com/vocdoni/census3/db/sqlc"
)

type TestDB struct {
	dir     string
	db      *sql.DB
	queries *queries.Queries
}

func StartTestDB(t *testing.T) *TestDB {
	c := qt.New(t)

	dir := t.TempDir()
	db, q, err := db.Init(dir)
	c.Assert(err, qt.IsNil)
	return &TestDB{dir, db, q}
}

func (testdb *TestDB) Close(t *testing.T) {
	c := qt.New(t)
	c.Assert(testdb.db.Close(), qt.IsNil)
	c.Assert(os.RemoveAll(testdb.dir), qt.IsNil)
}
