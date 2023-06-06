// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: metadata.sql

package queries

import (
	"context"
	"database/sql"
)

const chainID = `-- name: ChainID :one
SELECT chainID 
FROM metadata
LIMIT 1
`

func (q *Queries) ChainID(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, chainID)
	var chainid int64
	err := row.Scan(&chainid)
	return chainid, err
}

const setChainID = `-- name: SetChainID :execresult
INSERT INTO metadata (
    chainID
)
VALUES (
    ?
)
`

func (q *Queries) SetChainID(ctx context.Context, chainid int64) (sql.Result, error) {
	return q.db.ExecContext(ctx, setChainID, chainid)
}
