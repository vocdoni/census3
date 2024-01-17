// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: censuses.sql

package queries

import (
	"context"
	"database/sql"

	"github.com/vocdoni/census3/db/annotations"
)

const censusByID = `-- name: CensusByID :one
SELECT id, strategy_id, merkle_root, uri, size, weight, census_type, queue_id, accuracy FROM censuses
WHERE id = ?
LIMIT 1
`

func (q *Queries) CensusByID(ctx context.Context, id uint64) (Censuse, error) {
	row := q.db.QueryRowContext(ctx, censusByID, id)
	var i Censuse
	err := row.Scan(
		&i.ID,
		&i.StrategyID,
		&i.MerkleRoot,
		&i.Uri,
		&i.Size,
		&i.Weight,
		&i.CensusType,
		&i.QueueID,
		&i.Accuracy,
	)
	return i, err
}

const censusByStrategyID = `-- name: CensusByStrategyID :many
SELECT id, strategy_id, merkle_root, uri, size, weight, census_type, queue_id, accuracy FROM censuses
WHERE strategy_id = ?
`

func (q *Queries) CensusByStrategyID(ctx context.Context, strategyID uint64) ([]Censuse, error) {
	rows, err := q.db.QueryContext(ctx, censusByStrategyID, strategyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Censuse
	for rows.Next() {
		var i Censuse
		if err := rows.Scan(
			&i.ID,
			&i.StrategyID,
			&i.MerkleRoot,
			&i.Uri,
			&i.Size,
			&i.Weight,
			&i.CensusType,
			&i.QueueID,
			&i.Accuracy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createCensus = `-- name: CreateCensus :execresult
INSERT INTO censuses (
    id,
    strategy_id,
    merkle_root,
    uri,
    size, 
    weight,
    census_type,
    queue_id,
    accuracy
)
VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateCensusParams struct {
	ID         uint64
	StrategyID uint64
	MerkleRoot annotations.Hash
	Uri        sql.NullString
	Size       uint64
	Weight     sql.NullString
	CensusType uint64
	QueueID    string
	Accuracy   float64
}

func (q *Queries) CreateCensus(ctx context.Context, arg CreateCensusParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createCensus,
		arg.ID,
		arg.StrategyID,
		arg.MerkleRoot,
		arg.Uri,
		arg.Size,
		arg.Weight,
		arg.CensusType,
		arg.QueueID,
		arg.Accuracy,
	)
}
