// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: passports.sql

package queries

import (
	"context"
	"database/sql"
	"time"

	"github.com/vocdoni/census3/db/annotations"
)

const availableStamps = `-- name: AvailableStamps :many
SELECT DISTINCT(name) FROM stamps
`

func (q *Queries) AvailableStamps(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, availableStamps)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const deleteScore = `-- name: DeleteScore :execresult
DELETE FROM scores WHERE address = ?
`

func (q *Queries) DeleteScore(ctx context.Context, address annotations.Address) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteScore, address)
}

const deleteStampForAddress = `-- name: DeleteStampForAddress :execresult
DELETE FROM stamps WHERE address = ?
`

func (q *Queries) DeleteStampForAddress(ctx context.Context, address annotations.Address) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteStampForAddress, address)
}

const existsStamp = `-- name: ExistsStamp :one
SELECT EXISTS (
    SELECT name FROM (
        SELECT DISTINCT(name) FROM stamps
    ) WHERE name = ?
)
`

func (q *Queries) ExistsStamp(ctx context.Context, stamp string) (bool, error) {
	row := q.db.QueryRowContext(ctx, existsStamp, stamp)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getMetadata = `-- name: GetMetadata :one
SELECT value FROM metadata WHERE attr = ?
`

func (q *Queries) GetMetadata(ctx context.Context, attr string) (string, error) {
	row := q.db.QueryRowContext(ctx, getMetadata, attr)
	var value string
	err := row.Scan(&value)
	return value, err
}

const getScore = `-- name: GetScore :one
SELECT score FROM scores WHERE address = ?
`

func (q *Queries) GetScore(ctx context.Context, address annotations.Address) (annotations.BigInt, error) {
	row := q.db.QueryRowContext(ctx, getScore, address)
	var score annotations.BigInt
	err := row.Scan(&score)
	return score, err
}

const getScores = `-- name: GetScores :many
SELECT address, score FROM scores
`

type GetScoresRow struct {
	Address annotations.Address
	Score   annotations.BigInt
}

func (q *Queries) GetScores(ctx context.Context) ([]GetScoresRow, error) {
	rows, err := q.db.QueryContext(ctx, getScores)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetScoresRow
	for rows.Next() {
		var i GetScoresRow
		if err := rows.Scan(&i.Address, &i.Score); err != nil {
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

const getStampScoreForAddress = `-- name: GetStampScoreForAddress :one
SELECT score 
FROM stamps 
WHERE address = ? 
    AND name = ?
`

type GetStampScoreForAddressParams struct {
	Address annotations.Address
	Stamp   string
}

func (q *Queries) GetStampScoreForAddress(ctx context.Context, arg GetStampScoreForAddressParams) (annotations.BigInt, error) {
	row := q.db.QueryRowContext(ctx, getStampScoreForAddress, arg.Address, arg.Stamp)
	var score annotations.BigInt
	err := row.Scan(&score)
	return score, err
}

const getStampScores = `-- name: GetStampScores :many
SELECT address, score FROM stamps WHERE name = ?
`

type GetStampScoresRow struct {
	Address annotations.Address
	Score   annotations.BigInt
}

func (q *Queries) GetStampScores(ctx context.Context, name string) ([]GetStampScoresRow, error) {
	rows, err := q.db.QueryContext(ctx, getStampScores, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetStampScoresRow
	for rows.Next() {
		var i GetStampScoresRow
		if err := rows.Scan(&i.Address, &i.Score); err != nil {
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

const getStampsForAddress = `-- name: GetStampsForAddress :many
SELECT name, score FROM stamps WHERE address = ?
`

type GetStampsForAddressRow struct {
	Name  string
	Score annotations.BigInt
}

func (q *Queries) GetStampsForAddress(ctx context.Context, address annotations.Address) ([]GetStampsForAddressRow, error) {
	rows, err := q.db.QueryContext(ctx, getStampsForAddress, address)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetStampsForAddressRow
	for rows.Next() {
		var i GetStampsForAddressRow
		if err := rows.Scan(&i.Name, &i.Score); err != nil {
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

const newMetadata = `-- name: NewMetadata :exec
INSERT INTO metadata (attr, value) VALUES (?, ?)
`

type NewMetadataParams struct {
	Attr  string
	Value string
}

func (q *Queries) NewMetadata(ctx context.Context, arg NewMetadataParams) error {
	_, err := q.db.ExecContext(ctx, newMetadata, arg.Attr, arg.Value)
	return err
}

const newScore = `-- name: NewScore :execresult
INSERT INTO scores (address, score, date) VALUES (?, ?, ?)
`

type NewScoreParams struct {
	Address annotations.Address
	Score   annotations.BigInt
	Date    time.Time
}

func (q *Queries) NewScore(ctx context.Context, arg NewScoreParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, newScore, arg.Address, arg.Score, arg.Date)
}

const newStampScore = `-- name: NewStampScore :execresult
INSERT INTO stamps (address, name, score) VALUES (?, ?, ?)
`

type NewStampScoreParams struct {
	Address annotations.Address
	Name    string
	Score   annotations.BigInt
}

func (q *Queries) NewStampScore(ctx context.Context, arg NewStampScoreParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, newStampScore, arg.Address, arg.Name, arg.Score)
}

const scoreExists = `-- name: ScoreExists :one
SELECT EXISTS (
    SELECT address FROM scores WHERE address = ?
)
`

func (q *Queries) ScoreExists(ctx context.Context, address annotations.Address) (bool, error) {
	row := q.db.QueryRowContext(ctx, scoreExists, address)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const stampScoreExists = `-- name: StampScoreExists :one
SELECT EXISTS (
    SELECT address FROM stamps 
    WHERE address = ? 
        AND name = ?
)
`

type StampScoreExistsParams struct {
	Address annotations.Address
	Name    string
}

func (q *Queries) StampScoreExists(ctx context.Context, arg StampScoreExistsParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, stampScoreExists, arg.Address, arg.Name)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const stampTotalSupplyScores = `-- name: StampTotalSupplyScores :many
SELECT score FROM stamps WHERE name = ?
`

func (q *Queries) StampTotalSupplyScores(ctx context.Context, stamp string) ([]annotations.BigInt, error) {
	rows, err := q.db.QueryContext(ctx, stampTotalSupplyScores, stamp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []annotations.BigInt
	for rows.Next() {
		var score annotations.BigInt
		if err := rows.Scan(&score); err != nil {
			return nil, err
		}
		items = append(items, score)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const totalSupplyScores = `-- name: TotalSupplyScores :many
SELECT score FROM scores
`

func (q *Queries) TotalSupplyScores(ctx context.Context) ([]annotations.BigInt, error) {
	rows, err := q.db.QueryContext(ctx, totalSupplyScores)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []annotations.BigInt
	for rows.Next() {
		var score annotations.BigInt
		if err := rows.Scan(&score); err != nil {
			return nil, err
		}
		items = append(items, score)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMetadata = `-- name: UpdateMetadata :exec
UPDATE metadata
SET value = ?
WHERE attr = ?
`

type UpdateMetadataParams struct {
	Value string
	Attr  string
}

func (q *Queries) UpdateMetadata(ctx context.Context, arg UpdateMetadataParams) error {
	_, err := q.db.ExecContext(ctx, updateMetadata, arg.Value, arg.Attr)
	return err
}

const updateScore = `-- name: UpdateScore :execresult
UPDATE scores 
SET score = ?,
    date = ?
WHERE address = ?
`

type UpdateScoreParams struct {
	Score   annotations.BigInt
	Date    time.Time
	Address annotations.Address
}

func (q *Queries) UpdateScore(ctx context.Context, arg UpdateScoreParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateScore, arg.Score, arg.Date, arg.Address)
}

const updateStampScore = `-- name: UpdateStampScore :execresult
UPDATE stamps
SET score = ?
WHERE address = ? 
    AND name = ?
`

type UpdateStampScoreParams struct {
	Score   annotations.BigInt
	Address annotations.Address
	Name    string
}

func (q *Queries) UpdateStampScore(ctx context.Context, arg UpdateStampScoreParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateStampScore, arg.Score, arg.Address, arg.Name)
}
