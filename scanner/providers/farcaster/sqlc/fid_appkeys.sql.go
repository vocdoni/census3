// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: fid_appkeys.sql

package queries

import (
	"context"
	"database/sql"

	"github.com/vocdoni/census3/db/annotations"
)

const checkFidAppKeyExists = `-- name: CheckFidAppKeyExists :one
SELECT EXISTS (SELECT 1 FROM fid_appkeys WHERE fid = ? AND app_key = ?)
`

type CheckFidAppKeyExistsParams struct {
	Fid    uint64
	AppKey annotations.Bytes
}

func (q *Queries) CheckFidAppKeyExists(ctx context.Context, arg CheckFidAppKeyExistsParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkFidAppKeyExists, arg.Fid, arg.AppKey)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createFidAppKey = `-- name: CreateFidAppKey :execresult
INSERT INTO fid_appkeys (fid, app_key) VALUES (?, ?)
`

type CreateFidAppKeyParams struct {
	Fid    uint64
	AppKey annotations.Bytes
}

func (q *Queries) CreateFidAppKey(ctx context.Context, arg CreateFidAppKeyParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createFidAppKey, arg.Fid, arg.AppKey)
}

const deleteFidAppKey = `-- name: DeleteFidAppKey :execresult
DELETE FROM fid_appkeys WHERE fid = ? AND app_key = ?
`

type DeleteFidAppKeyParams struct {
	Fid    uint64
	AppKey annotations.Bytes
}

func (q *Queries) DeleteFidAppKey(ctx context.Context, arg DeleteFidAppKeyParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteFidAppKey, arg.Fid, arg.AppKey)
}

const getFidAppKeys = `-- name: GetFidAppKeys :many
SELECT app_key FROM fid_appkeys WHERE fid = ?
`

func (q *Queries) GetFidAppKeys(ctx context.Context, fid uint64) ([]annotations.Bytes, error) {
	rows, err := q.db.QueryContext(ctx, getFidAppKeys, fid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []annotations.Bytes
	for rows.Next() {
		var app_key annotations.Bytes
		if err := rows.Scan(&app_key); err != nil {
			return nil, err
		}
		items = append(items, app_key)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFidByAppKey = `-- name: GetFidByAppKey :many
SELECT fid FROM fid_appkeys WHERE app_key = ?
`

func (q *Queries) GetFidByAppKey(ctx context.Context, appKey annotations.Bytes) ([]uint64, error) {
	rows, err := q.db.QueryContext(ctx, getFidByAppKey, appKey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []uint64
	for rows.Next() {
		var fid uint64
		if err := rows.Scan(&fid); err != nil {
			return nil, err
		}
		items = append(items, fid)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAppKeys = `-- name: ListAppKeys :many
SELECT app_key FROM fid_appkeys ORDER BY app_key ASC
`

func (q *Queries) ListAppKeys(ctx context.Context) ([]annotations.Bytes, error) {
	rows, err := q.db.QueryContext(ctx, listAppKeys)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []annotations.Bytes
	for rows.Next() {
		var app_key annotations.Bytes
		if err := rows.Scan(&app_key); err != nil {
			return nil, err
		}
		items = append(items, app_key)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}