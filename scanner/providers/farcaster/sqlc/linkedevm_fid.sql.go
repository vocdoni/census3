// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: linkedevm_fid.sql

package queries

import (
	"context"
	"database/sql"

	"github.com/vocdoni/census3/db/annotations"
)

const checkLinkedEVMForAny = `-- name: CheckLinkedEVMForAny :one
SELECT EXISTS(SELECT 1 FROM linkedevm_fid WHERE evm_address = ?)
`

func (q *Queries) CheckLinkedEVMForAny(ctx context.Context, evmAddress []byte) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkLinkedEVMForAny, evmAddress)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createLinkedEVMFID = `-- name: CreateLinkedEVMFID :execresult
INSERT INTO linkedevm_fid (fid, evm_address) VALUES (?, ?)
`

type CreateLinkedEVMFIDParams struct {
	Fid        uint64
	EvmAddress []byte
}

func (q *Queries) CreateLinkedEVMFID(ctx context.Context, arg CreateLinkedEVMFIDParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createLinkedEVMFID, arg.Fid, arg.EvmAddress)
}

const deleteLinkedEVMFID = `-- name: DeleteLinkedEVMFID :execresult
DELETE FROM linkedevm_fid WHERE fid = ? AND evm_address = ?
`

type DeleteLinkedEVMFIDParams struct {
	Fid        uint64
	EvmAddress []byte
}

func (q *Queries) DeleteLinkedEVMFID(ctx context.Context, arg DeleteLinkedEVMFIDParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteLinkedEVMFID, arg.Fid, arg.EvmAddress)
}

const getLinkedEVMFID = `-- name: GetLinkedEVMFID :many
SELECT evm_address FROM linkedevm_fid WHERE fid = ?
`

func (q *Queries) GetLinkedEVMFID(ctx context.Context, fid uint64) ([][]byte, error) {
	rows, err := q.db.QueryContext(ctx, getLinkedEVMFID, fid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items [][]byte
	for rows.Next() {
		var evm_address []byte
		if err := rows.Scan(&evm_address); err != nil {
			return nil, err
		}
		items = append(items, evm_address)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsersNotLinkedEVM = `-- name: ListUsersNotLinkedEVM :many
SELECT fid, username, signer, custody_address, app_keys, recovery_address, linked_evm FROM users WHERE fid NOT IN (SELECT fid FROM linkedevm_fid)
`

func (q *Queries) ListUsersNotLinkedEVM(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsersNotLinkedEVM)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.Fid,
			&i.Username,
			&i.Signer,
			&i.CustodyAddress,
			&i.AppKeys,
			&i.RecoveryAddress,
			&i.LinkedEvm,
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

const listUsersWithLinkedEVM = `-- name: ListUsersWithLinkedEVM :many
SELECT u.fid, u.signer, l.evm_address FROM users u
JOIN linkedevm_fid l ON u.fid = l.fid
`

type ListUsersWithLinkedEVMRow struct {
	Fid        uint64
	Signer     annotations.Bytes
	EvmAddress []byte
}

func (q *Queries) ListUsersWithLinkedEVM(ctx context.Context) ([]ListUsersWithLinkedEVMRow, error) {
	rows, err := q.db.QueryContext(ctx, listUsersWithLinkedEVM)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersWithLinkedEVMRow
	for rows.Next() {
		var i ListUsersWithLinkedEVMRow
		if err := rows.Scan(&i.Fid, &i.Signer, &i.EvmAddress); err != nil {
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
