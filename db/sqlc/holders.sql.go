// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: holders.sql

package queries

import (
	"context"
	"database/sql"

	"github.com/vocdoni/census3/db"
)

const createHolder = `-- name: CreateHolder :execresult
INSERT INTO Holders (id)
VALUES (?)
`

func (q *Queries) CreateHolder(ctx context.Context, id db.Address) (sql.Result, error) {
	return q.db.ExecContext(ctx, createHolder, id)
}

const createTokenHolder = `-- name: CreateTokenHolder :execresult
INSERT INTO TokenHolders (
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
	TokenID  db.Address
	HolderID db.Address
	Balance  db.BigInt
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
DELETE FROM Holders
WHERE id = ?
`

func (q *Queries) DeleteHolder(ctx context.Context, id db.Address) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteHolder, id)
}

const deleteTokenHolder = `-- name: DeleteTokenHolder :execresult
DELETE FROM TokenHolders
WHERE token_id = ? AND holder_id = ? AND block_id = ?
`

type DeleteTokenHolderParams struct {
	TokenID  db.Address
	HolderID db.Address
	BlockID  int64
}

func (q *Queries) DeleteTokenHolder(ctx context.Context, arg DeleteTokenHolderParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteTokenHolder, arg.TokenID, arg.HolderID, arg.BlockID)
}

const holderByID = `-- name: HolderByID :one
SELECT id FROM Holders
WHERE id = ?
LIMIT 1
`

func (q *Queries) HolderByID(ctx context.Context, id db.Address) (db.Address, error) {
	row := q.db.QueryRowContext(ctx, holderByID, id)
	err := row.Scan(&id)
	return id, err
}

const lastBlockByTokenIDAndHolderID = `-- name: LastBlockByTokenIDAndHolderID :one
SELECT block_id 
FROM TokenHolders
WHERE token_id = ?
ORDER BY block_id DESC
LIMIT 1
`

func (q *Queries) LastBlockByTokenIDAndHolderID(ctx context.Context, tokenID db.Address) (int64, error) {
	row := q.db.QueryRowContext(ctx, lastBlockByTokenIDAndHolderID, tokenID)
	var block_id int64
	err := row.Scan(&block_id)
	return block_id, err
}

const paginatedHolders = `-- name: PaginatedHolders :many
SELECT id FROM Holders
ORDER BY id
LIMIT ? OFFSET ?
`

type PaginatedHoldersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) PaginatedHolders(ctx context.Context, arg PaginatedHoldersParams) ([]db.Address, error) {
	rows, err := q.db.QueryContext(ctx, paginatedHolders, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []db.Address
	for rows.Next() {
		var id db.Address
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

const tokenHolderByTokenIDAndBlockIDAndHolderID = `-- name: TokenHolderByTokenIDAndBlockIDAndHolderID :one
SELECT holders.id, TokenHolders.balance
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = ? AND TokenHolders.holder_id = ? AND TokenHolders.block_id = ?
`

type TokenHolderByTokenIDAndBlockIDAndHolderIDParams struct {
	TokenID  db.Address
	HolderID db.Address
	BlockID  int64
}

type TokenHolderByTokenIDAndBlockIDAndHolderIDRow struct {
	ID      db.Address
	Balance db.BigInt
}

func (q *Queries) TokenHolderByTokenIDAndBlockIDAndHolderID(ctx context.Context, arg TokenHolderByTokenIDAndBlockIDAndHolderIDParams) (TokenHolderByTokenIDAndBlockIDAndHolderIDRow, error) {
	row := q.db.QueryRowContext(ctx, tokenHolderByTokenIDAndBlockIDAndHolderID, arg.TokenID, arg.HolderID, arg.BlockID)
	var i TokenHolderByTokenIDAndBlockIDAndHolderIDRow
	err := row.Scan(&i.ID, &i.Balance)
	return i, err
}

const tokenHolderByTokenIDAndHolderID = `-- name: TokenHolderByTokenIDAndHolderID :one
SELECT holders.id, TokenHolders.balance
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = ? AND TokenHolders.holder_id = ?
`

type TokenHolderByTokenIDAndHolderIDParams struct {
	TokenID  db.Address
	HolderID db.Address
}

type TokenHolderByTokenIDAndHolderIDRow struct {
	ID      db.Address
	Balance db.BigInt
}

func (q *Queries) TokenHolderByTokenIDAndHolderID(ctx context.Context, arg TokenHolderByTokenIDAndHolderIDParams) (TokenHolderByTokenIDAndHolderIDRow, error) {
	row := q.db.QueryRowContext(ctx, tokenHolderByTokenIDAndHolderID, arg.TokenID, arg.HolderID)
	var i TokenHolderByTokenIDAndHolderIDRow
	err := row.Scan(&i.ID, &i.Balance)
	return i, err
}

const tokenHoldersByTokenID = `-- name: TokenHoldersByTokenID :many
SELECT holders.id
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = ?
LIMIT ? OFFSET ?
`

type TokenHoldersByTokenIDParams struct {
	TokenID db.Address
	Limit   int32
	Offset  int32
}

func (q *Queries) TokenHoldersByTokenID(ctx context.Context, arg TokenHoldersByTokenIDParams) ([]db.Address, error) {
	rows, err := q.db.QueryContext(ctx, tokenHoldersByTokenID, arg.TokenID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []db.Address
	for rows.Next() {
		var id db.Address
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

const tokenHoldersByTokenIDAndBlockID = `-- name: TokenHoldersByTokenIDAndBlockID :many
SELECT holders.id
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = ? AND TokenHolders.block_id = ?
LIMIT ? OFFSET ?
`

type TokenHoldersByTokenIDAndBlockIDParams struct {
	TokenID db.Address
	BlockID int64
	Limit   int32
	Offset  int32
}

func (q *Queries) TokenHoldersByTokenIDAndBlockID(ctx context.Context, arg TokenHoldersByTokenIDAndBlockIDParams) ([]db.Address, error) {
	rows, err := q.db.QueryContext(ctx, tokenHoldersByTokenIDAndBlockID,
		arg.TokenID,
		arg.BlockID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []db.Address
	for rows.Next() {
		var id db.Address
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
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = ? AND TokenHolders.balance >= ? AND TokenHolders.block_id = ?
LIMIT ? OFFSET ?
`

type TokenHoldersByTokenIDAndBlockIDAndMinBalanceParams struct {
	TokenID db.Address
	Balance db.BigInt
	BlockID int64
	Limit   int32
	Offset  int32
}

func (q *Queries) TokenHoldersByTokenIDAndBlockIDAndMinBalance(ctx context.Context, arg TokenHoldersByTokenIDAndBlockIDAndMinBalanceParams) ([]db.Address, error) {
	rows, err := q.db.QueryContext(ctx, tokenHoldersByTokenIDAndBlockIDAndMinBalance,
		arg.TokenID,
		arg.Balance,
		arg.BlockID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []db.Address
	for rows.Next() {
		var id db.Address
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
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = ? AND TokenHolders.balance >= ?
LIMIT ? OFFSET ?
`

type TokenHoldersByTokenIDAndMinBalanceParams struct {
	TokenID db.Address
	Balance db.BigInt
	Limit   int32
	Offset  int32
}

func (q *Queries) TokenHoldersByTokenIDAndMinBalance(ctx context.Context, arg TokenHoldersByTokenIDAndMinBalanceParams) ([]db.Address, error) {
	rows, err := q.db.QueryContext(ctx, tokenHoldersByTokenIDAndMinBalance,
		arg.TokenID,
		arg.Balance,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []db.Address
	for rows.Next() {
		var id db.Address
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

const tokenHoldersPaginated = `-- name: TokenHoldersPaginated :many
SELECT token_id, holder_id, balance, block_id FROM TokenHolders
ORDER BY token_id, holder_id, block_id
LIMIT ? OFFSET ?
`

type TokenHoldersPaginatedParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) TokenHoldersPaginated(ctx context.Context, arg TokenHoldersPaginatedParams) ([]Tokenholder, error) {
	rows, err := q.db.QueryContext(ctx, tokenHoldersPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tokenholder
	for rows.Next() {
		var i Tokenholder
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

const tokensByHolderID = `-- name: TokensByHolderID :many
SELECT tokens.id, tokens.name, tokens.symbol, tokens.decimals, tokens.total_supply, tokens.creation_block, tokens.type_id
FROM Tokens
JOIN TokenHolders ON Tokens.id = TokenHolders.token_id
WHERE TokenHolders.holder_id = ?
LIMIT ? OFFSET ?
`

type TokensByHolderIDParams struct {
	HolderID db.Address
	Limit    int32
	Offset   int32
}

func (q *Queries) TokensByHolderID(ctx context.Context, arg TokensByHolderIDParams) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, tokensByHolderID, arg.HolderID, arg.Limit, arg.Offset)
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
SELECT tokens.id, tokens.name, tokens.symbol, tokens.decimals, tokens.total_supply, tokens.creation_block, tokens.type_id
FROM Tokens
JOIN TokenHolders ON Tokens.id = TokenHolders.token_id
WHERE TokenHolders.holder_id = ? AND TokenHolders.block_id = ?
LIMIT ? OFFSET ?
`

type TokensByHolderIDAndBlockIDParams struct {
	HolderID db.Address
	BlockID  int64
	Limit    int32
	Offset   int32
}

func (q *Queries) TokensByHolderIDAndBlockID(ctx context.Context, arg TokensByHolderIDAndBlockIDParams) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, tokensByHolderIDAndBlockID,
		arg.HolderID,
		arg.BlockID,
		arg.Limit,
		arg.Offset,
	)
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

const updateTokenHolder = `-- name: UpdateTokenHolder :execresult
UPDATE TokenHolders
SET balance = ?,
    block_id = ?
WHERE token_id = ? AND holder_id = ? AND block_id = ?
`

type UpdateTokenHolderParams struct {
	Balance  db.BigInt
	BlockID  int64
	TokenID  db.Address
	HolderID db.Address
}

func (q *Queries) UpdateTokenHolder(ctx context.Context, arg UpdateTokenHolderParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateTokenHolder,
		arg.Balance,
		arg.BlockID,
		arg.TokenID,
		arg.HolderID,
		arg.BlockID,
	)
}
