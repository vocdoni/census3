// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: holders.sql

package queries

import (
	"context"
	"database/sql"

	"github.com/vocdoni/census3/db/annotations"
)

const countTokenHoldersByTokenID = `-- name: CountTokenHoldersByTokenID :one
SELECT COUNT(holder_id) 
FROM token_holders
WHERE token_id = ?
`

func (q *Queries) CountTokenHoldersByTokenID(ctx context.Context, tokenID []byte) (int64, error) {
	row := q.db.QueryRowContext(ctx, countTokenHoldersByTokenID, tokenID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createHolder = `-- name: CreateHolder :execresult
INSERT INTO holders (id)
VALUES (?)
`

func (q *Queries) CreateHolder(ctx context.Context, id annotations.Address) (sql.Result, error) {
	return q.db.ExecContext(ctx, createHolder, id)
}

const createTokenHolder = `-- name: CreateTokenHolder :execresult
INSERT INTO token_holders (
    token_id,
    holder_id,
    balance,
    block_id
)
VALUES (
    ?, ?, ?, ?
)
`

type CreateTokenHolderParams struct {
	TokenID  []byte
	HolderID []byte
	Balance  []byte
	BlockID  int64
}

func (q *Queries) CreateTokenHolder(ctx context.Context, arg CreateTokenHolderParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createTokenHolder,
		arg.TokenID,
		arg.HolderID,
		arg.Balance,
		arg.BlockID,
	)
}

const deleteHolder = `-- name: DeleteHolder :execresult
DELETE FROM holders
WHERE id = ?
`

func (q *Queries) DeleteHolder(ctx context.Context, id annotations.Address) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteHolder, id)
}

const deleteTokenHolder = `-- name: DeleteTokenHolder :execresult
DELETE FROM token_holders
WHERE token_id = ? AND holder_id = ?
`

type DeleteTokenHolderParams struct {
	TokenID  []byte
	HolderID []byte
}

func (q *Queries) DeleteTokenHolder(ctx context.Context, arg DeleteTokenHolderParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteTokenHolder, arg.TokenID, arg.HolderID)
}

const holderByID = `-- name: HolderByID :one
SELECT id FROM holders
WHERE id = ?
LIMIT 1
`

func (q *Queries) HolderByID(ctx context.Context, id annotations.Address) (annotations.Address, error) {
	row := q.db.QueryRowContext(ctx, holderByID, id)
	err := row.Scan(&id)
	return id, err
}

const lastBlockByTokenID = `-- name: LastBlockByTokenID :one
SELECT block_id 
FROM token_holders
WHERE token_id = ?
ORDER BY block_id DESC
LIMIT 1
`

func (q *Queries) LastBlockByTokenID(ctx context.Context, tokenID []byte) (int64, error) {
	row := q.db.QueryRowContext(ctx, lastBlockByTokenID, tokenID)
	var block_id int64
	err := row.Scan(&block_id)
	return block_id, err
}

const listHolders = `-- name: ListHolders :many
SELECT id FROM holders
ORDER BY id
`

func (q *Queries) ListHolders(ctx context.Context) ([]annotations.Address, error) {
	rows, err := q.db.QueryContext(ctx, listHolders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []annotations.Address
	for rows.Next() {
		var id annotations.Address
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTokenHolders = `-- name: ListTokenHolders :many
SELECT token_id, holder_id, balance, block_id FROM token_holders
ORDER BY token_id, holder_id, block_id
`

func (q *Queries) ListTokenHolders(ctx context.Context) ([]TokenHolder, error) {
	rows, err := q.db.QueryContext(ctx, listTokenHolders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TokenHolder
	for rows.Next() {
		var i TokenHolder
		if err := rows.Scan(
			&i.TokenID,
			&i.HolderID,
			&i.Balance,
			&i.BlockID,
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

const tokenHolderByTokenIDAndBlockIDAndHolderID = `-- name: TokenHolderByTokenIDAndBlockIDAndHolderID :one
SELECT holders.id, token_holders.balance
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = ? AND token_holders.holder_id = ? AND token_holders.block_id = ?
`

type TokenHolderByTokenIDAndBlockIDAndHolderIDParams struct {
	TokenID  []byte
	HolderID []byte
	BlockID  int64
}

type TokenHolderByTokenIDAndBlockIDAndHolderIDRow struct {
	ID      annotations.Address
	Balance []byte
}

func (q *Queries) TokenHolderByTokenIDAndBlockIDAndHolderID(ctx context.Context, arg TokenHolderByTokenIDAndBlockIDAndHolderIDParams) (TokenHolderByTokenIDAndBlockIDAndHolderIDRow, error) {
	row := q.db.QueryRowContext(ctx, tokenHolderByTokenIDAndBlockIDAndHolderID, arg.TokenID, arg.HolderID, arg.BlockID)
	var i TokenHolderByTokenIDAndBlockIDAndHolderIDRow
	err := row.Scan(&i.ID, &i.Balance)
	return i, err
}

const tokenHolderByTokenIDAndHolderID = `-- name: TokenHolderByTokenIDAndHolderID :one
SELECT holders.id, token_holders.token_id, token_holders.holder_id, token_holders.balance, token_holders.block_id
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = ? AND token_holders.holder_id = ?
`

type TokenHolderByTokenIDAndHolderIDParams struct {
	TokenID  []byte
	HolderID []byte
}

type TokenHolderByTokenIDAndHolderIDRow struct {
	ID       annotations.Address
	TokenID  []byte
	HolderID []byte
	Balance  []byte
	BlockID  int64
}

func (q *Queries) TokenHolderByTokenIDAndHolderID(ctx context.Context, arg TokenHolderByTokenIDAndHolderIDParams) (TokenHolderByTokenIDAndHolderIDRow, error) {
	row := q.db.QueryRowContext(ctx, tokenHolderByTokenIDAndHolderID, arg.TokenID, arg.HolderID)
	var i TokenHolderByTokenIDAndHolderIDRow
	err := row.Scan(
		&i.ID,
		&i.TokenID,
		&i.HolderID,
		&i.Balance,
		&i.BlockID,
	)
	return i, err
}

const tokenHoldersByTokenID = `-- name: TokenHoldersByTokenID :many
SELECT holders.id, token_holders.balance
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = ?
`

type TokenHoldersByTokenIDRow struct {
	ID      annotations.Address
	Balance []byte
}

func (q *Queries) TokenHoldersByTokenID(ctx context.Context, tokenID []byte) ([]TokenHoldersByTokenIDRow, error) {
	rows, err := q.db.QueryContext(ctx, tokenHoldersByTokenID, tokenID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TokenHoldersByTokenIDRow
	for rows.Next() {
		var i TokenHoldersByTokenIDRow
		if err := rows.Scan(&i.ID, &i.Balance); err != nil {
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

const tokenHoldersByTokenIDAndBlockID = `-- name: TokenHoldersByTokenIDAndBlockID :many
SELECT holders.id
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = ? AND token_holders.block_id = ?
`

type TokenHoldersByTokenIDAndBlockIDParams struct {
	TokenID []byte
	BlockID int64
}

func (q *Queries) TokenHoldersByTokenIDAndBlockID(ctx context.Context, arg TokenHoldersByTokenIDAndBlockIDParams) ([]annotations.Address, error) {
	rows, err := q.db.QueryContext(ctx, tokenHoldersByTokenIDAndBlockID, arg.TokenID, arg.BlockID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []annotations.Address
	for rows.Next() {
		var id annotations.Address
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const tokenHoldersByTokenIDAndBlockIDAndMinBalance = `-- name: TokenHoldersByTokenIDAndBlockIDAndMinBalance :many
SELECT holders.id
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = ? AND token_holders.balance >= ? AND token_holders.block_id = ?
`

type TokenHoldersByTokenIDAndBlockIDAndMinBalanceParams struct {
	TokenID []byte
	Balance []byte
	BlockID int64
}

func (q *Queries) TokenHoldersByTokenIDAndBlockIDAndMinBalance(ctx context.Context, arg TokenHoldersByTokenIDAndBlockIDAndMinBalanceParams) ([]annotations.Address, error) {
	rows, err := q.db.QueryContext(ctx, tokenHoldersByTokenIDAndBlockIDAndMinBalance, arg.TokenID, arg.Balance, arg.BlockID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []annotations.Address
	for rows.Next() {
		var id annotations.Address
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const tokenHoldersByTokenIDAndMinBalance = `-- name: TokenHoldersByTokenIDAndMinBalance :many
SELECT holders.id
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = ? AND token_holders.balance >= ?
`

type TokenHoldersByTokenIDAndMinBalanceParams struct {
	TokenID []byte
	Balance []byte
}

func (q *Queries) TokenHoldersByTokenIDAndMinBalance(ctx context.Context, arg TokenHoldersByTokenIDAndMinBalanceParams) ([]annotations.Address, error) {
	rows, err := q.db.QueryContext(ctx, tokenHoldersByTokenIDAndMinBalance, arg.TokenID, arg.Balance)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []annotations.Address
	for rows.Next() {
		var id annotations.Address
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const tokensByHolderID = `-- name: TokensByHolderID :many
SELECT tokens.id, tokens.name, tokens.symbol, tokens.decimals, tokens.total_supply, tokens.creation_block, tokens.type_id, tokens.synced, tokens.tag
FROM Tokens
JOIN token_holders ON tokens.id = token_holders.token_id
WHERE token_holders.holder_id = ?
`

func (q *Queries) TokensByHolderID(ctx context.Context, holderID []byte) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, tokensByHolderID, holderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Token
	for rows.Next() {
		var i Token
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Symbol,
			&i.Decimals,
			&i.TotalSupply,
			&i.CreationBlock,
			&i.TypeID,
			&i.Synced,
			&i.Tag,
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

const tokensByHolderIDAndBlockID = `-- name: TokensByHolderIDAndBlockID :many
SELECT tokens.id, tokens.name, tokens.symbol, tokens.decimals, tokens.total_supply, tokens.creation_block, tokens.type_id, tokens.synced, tokens.tag
FROM Tokens
JOIN token_holders ON tokens.id = token_holders.token_id
WHERE token_holders.holder_id = ? AND token_holders.block_id = ?
`

type TokensByHolderIDAndBlockIDParams struct {
	HolderID []byte
	BlockID  int64
}

func (q *Queries) TokensByHolderIDAndBlockID(ctx context.Context, arg TokensByHolderIDAndBlockIDParams) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, tokensByHolderIDAndBlockID, arg.HolderID, arg.BlockID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Token
	for rows.Next() {
		var i Token
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Symbol,
			&i.Decimals,
			&i.TotalSupply,
			&i.CreationBlock,
			&i.TypeID,
			&i.Synced,
			&i.Tag,
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

const updateTokenHolderBalance = `-- name: UpdateTokenHolderBalance :execresult
UPDATE token_holders
SET balance = ?,
    block_id = ?
WHERE token_id = ? AND holder_id = ? AND block_id = ?
`

type UpdateTokenHolderBalanceParams struct {
	Balance    []byte
	NewBlockID int64
	TokenID    []byte
	HolderID   []byte
	BlockID    int64
}

func (q *Queries) UpdateTokenHolderBalance(ctx context.Context, arg UpdateTokenHolderBalanceParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateTokenHolderBalance,
		arg.Balance,
		arg.NewBlockID,
		arg.TokenID,
		arg.HolderID,
		arg.BlockID,
	)
}
