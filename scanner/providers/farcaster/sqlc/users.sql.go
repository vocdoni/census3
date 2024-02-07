// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: users.sql

package queries

import (
	"context"
	"database/sql"

	"github.com/vocdoni/census3/db/annotations"
)

const countUserAppKeys = `-- name: CountUserAppKeys :one
SELECT COUNT(app_keys) FROM users WHERE fid = ?
`

func (q *Queries) CountUserAppKeys(ctx context.Context, fid uint64) (int64, error) {
	row := q.db.QueryRowContext(ctx, countUserAppKeys, fid)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countUsers = `-- name: CountUsers :one
SELECT COUNT(*) FROM users
`

func (q *Queries) CountUsers(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countUsers)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (
    fid,
    username,
    signer,
    custody_address,
    app_keys,
    recovery_address)
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateUserParams struct {
	Fid             uint64
	Username        string
	Signer          annotations.Bytes
	CustodyAddress  annotations.Address
	AppKeys         sql.NullString
	RecoveryAddress annotations.Address
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.Fid,
		arg.Username,
		arg.Signer,
		arg.CustodyAddress,
		arg.AppKeys,
		arg.RecoveryAddress,
	)
}

const deleteUser = `-- name: DeleteUser :execresult
DELETE FROM users WHERE fid = ?
`

func (q *Queries) DeleteUser(ctx context.Context, fid uint64) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteUser, fid)
}

const getUserByFID = `-- name: GetUserByFID :one
SELECT fid, username, signer, custody_address, app_keys, recovery_address FROM users WHERE fid = ?
`

func (q *Queries) GetUserByFID(ctx context.Context, fid uint64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByFID, fid)
	var i User
	err := row.Scan(
		&i.Fid,
		&i.Username,
		&i.Signer,
		&i.CustodyAddress,
		&i.AppKeys,
		&i.RecoveryAddress,
	)
	return i, err
}

const getUserBySigner = `-- name: GetUserBySigner :one
SELECT fid, username, signer, custody_address, app_keys, recovery_address FROM users WHERE signer = ?
`

func (q *Queries) GetUserBySigner(ctx context.Context, signer annotations.Bytes) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserBySigner, signer)
	var i User
	err := row.Scan(
		&i.Fid,
		&i.Username,
		&i.Signer,
		&i.CustodyAddress,
		&i.AppKeys,
		&i.RecoveryAddress,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT fid, username, signer, custody_address, app_keys, recovery_address FROM users WHERE username = ?
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.Fid,
		&i.Username,
		&i.Signer,
		&i.CustodyAddress,
		&i.AppKeys,
		&i.RecoveryAddress,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT fid, username, signer, custody_address, app_keys, recovery_address FROM users ORDER BY fid ASC
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
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

const updateCustodyAddress = `-- name: UpdateCustodyAddress :execresult
UPDATE users
SET custody_address = ?
WHERE fid = ?
`

type UpdateCustodyAddressParams struct {
	CustodyAddress annotations.Address
	Fid            uint64
}

func (q *Queries) UpdateCustodyAddress(ctx context.Context, arg UpdateCustodyAddressParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateCustodyAddress, arg.CustodyAddress, arg.Fid)
}

const updateUser = `-- name: UpdateUser :execresult
UPDATE users 
SET username = ?,
    signer = ?,
    custody_address = ?,
    app_keys = ?,
    recovery_address = ?
WHERE fid = ?
`

type UpdateUserParams struct {
	Username        string
	Signer          annotations.Bytes
	CustodyAddress  annotations.Address
	AppKeys         sql.NullString
	RecoveryAddress annotations.Address
	Fid             uint64
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUser,
		arg.Username,
		arg.Signer,
		arg.CustodyAddress,
		arg.AppKeys,
		arg.RecoveryAddress,
		arg.Fid,
	)
}

const updateUserAppKeys = `-- name: UpdateUserAppKeys :execresult
UPDATE users
SET app_keys = ?
WHERE fid = ?
`

type UpdateUserAppKeysParams struct {
	AppKeys sql.NullString
	Fid     uint64
}

func (q *Queries) UpdateUserAppKeys(ctx context.Context, arg UpdateUserAppKeysParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUserAppKeys, arg.AppKeys, arg.Fid)
}

const updateUserRecoveryAddress = `-- name: UpdateUserRecoveryAddress :execresult
UPDATE users
SET recovery_address = ?
WHERE fid = ?
`

type UpdateUserRecoveryAddressParams struct {
	RecoveryAddress annotations.Address
	Fid             uint64
}

func (q *Queries) UpdateUserRecoveryAddress(ctx context.Context, arg UpdateUserRecoveryAddressParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUserRecoveryAddress, arg.RecoveryAddress, arg.Fid)
}

const updateUsername = `-- name: UpdateUsername :execresult
UPDATE users
SET username = ?
WHERE fid = ?
`

type UpdateUsernameParams struct {
	Username string
	Fid      uint64
}

func (q *Queries) UpdateUsername(ctx context.Context, arg UpdateUsernameParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUsername, arg.Username, arg.Fid)
}
