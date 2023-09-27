// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: tokens.sql

package queries

import (
	"context"
	"database/sql"

	"github.com/vocdoni/census3/db/annotations"
)

const createToken = `-- name: CreateToken :execresult
INSERT INTO tokens (
    id,
    name,
    symbol,
    decimals,
    total_supply,
    creation_block,
    type_id,
    synced,
    tags,
    chain_id
)
VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateTokenParams struct {
	ID            annotations.Address
	Name          sql.NullString
	Symbol        sql.NullString
	Decimals      uint64
	TotalSupply   annotations.BigInt
	CreationBlock sql.NullInt64
	TypeID        uint64
	Synced        bool
	Tags          sql.NullString
	ChainID       uint64
}

func (q *Queries) CreateToken(ctx context.Context, arg CreateTokenParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createToken,
		arg.ID,
		arg.Name,
		arg.Symbol,
		arg.Decimals,
		arg.TotalSupply,
		arg.CreationBlock,
		arg.TypeID,
		arg.Synced,
		arg.Tags,
		arg.ChainID,
	)
}

const existsToken = `-- name: ExistsToken :one
SELECT EXISTS 
    (SELECT id 
    FROM tokens
    WHERE id = ?)
`

func (q *Queries) ExistsToken(ctx context.Context, id annotations.Address) (bool, error) {
	row := q.db.QueryRowContext(ctx, existsToken, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const existsTokenByChainID = `-- name: ExistsTokenByChainID :one
SELECT EXISTS 
    (SELECT id 
    FROM tokens
    WHERE id = ? AND chain_id = ?)
`

type ExistsTokenByChainIDParams struct {
	ID      annotations.Address
	ChainID uint64
}

func (q *Queries) ExistsTokenByChainID(ctx context.Context, arg ExistsTokenByChainIDParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, existsTokenByChainID, arg.ID, arg.ChainID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const listTokens = `-- name: ListTokens :many
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id FROM tokens
ORDER BY type_id, name
`

func (q *Queries) ListTokens(ctx context.Context) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, listTokens)
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
			&i.Tags,
			&i.ChainID,
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

const tokenByID = `-- name: TokenByID :one
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id FROM tokens
WHERE id = ?
LIMIT 1
`

func (q *Queries) TokenByID(ctx context.Context, id annotations.Address) (Token, error) {
	row := q.db.QueryRowContext(ctx, tokenByID, id)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Symbol,
		&i.Decimals,
		&i.TotalSupply,
		&i.CreationBlock,
		&i.TypeID,
		&i.Synced,
		&i.Tags,
		&i.ChainID,
	)
	return i, err
}

const tokensByStrategyID = `-- name: TokensByStrategyID :many
SELECT t.id, t.name, t.symbol, t.decimals, t.total_supply, t.creation_block, t.type_id, t.synced, t.tags, t.chain_id, st.strategy_id, st.token_id, st.chain_id, st.min_balance FROM tokens t
JOIN strategy_tokens st ON st.token_id = t.id
WHERE st.strategy_id = ?
ORDER BY t.name
`

type TokensByStrategyIDRow struct {
	ID            annotations.Address
	Name          sql.NullString
	Symbol        sql.NullString
	Decimals      uint64
	TotalSupply   annotations.BigInt
	CreationBlock sql.NullInt64
	TypeID        uint64
	Synced        bool
	Tags          sql.NullString
	ChainID       uint64
	StrategyID    uint64
	TokenID       []byte
	ChainID_2     uint64
	MinBalance    []byte
}

func (q *Queries) TokensByStrategyID(ctx context.Context, strategyID uint64) ([]TokensByStrategyIDRow, error) {
	rows, err := q.db.QueryContext(ctx, tokensByStrategyID, strategyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TokensByStrategyIDRow
	for rows.Next() {
		var i TokensByStrategyIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Symbol,
			&i.Decimals,
			&i.TotalSupply,
			&i.CreationBlock,
			&i.TypeID,
			&i.Synced,
			&i.Tags,
			&i.ChainID,
			&i.StrategyID,
			&i.TokenID,
			&i.ChainID_2,
			&i.MinBalance,
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

const updateTokenCreationBlock = `-- name: UpdateTokenCreationBlock :execresult
UPDATE tokens
SET creation_block = ?
WHERE id = ?
`

type UpdateTokenCreationBlockParams struct {
	CreationBlock sql.NullInt64
	ID            annotations.Address
}

func (q *Queries) UpdateTokenCreationBlock(ctx context.Context, arg UpdateTokenCreationBlockParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateTokenCreationBlock, arg.CreationBlock, arg.ID)
}

const updateTokenStatus = `-- name: UpdateTokenStatus :execresult
UPDATE tokens
SET synced = ?
WHERE id = ?
`

type UpdateTokenStatusParams struct {
	Synced bool
	ID     annotations.Address
}

func (q *Queries) UpdateTokenStatus(ctx context.Context, arg UpdateTokenStatusParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateTokenStatus, arg.Synced, arg.ID)
}
