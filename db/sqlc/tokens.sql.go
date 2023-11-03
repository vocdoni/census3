// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
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
    chain_id,
    chain_address
)
VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateTokenParams struct {
	ID            annotations.Address
	Name          string
	Symbol        string
	Decimals      uint64
	TotalSupply   annotations.BigInt
	CreationBlock int64
	TypeID        uint64
	Synced        bool
	Tags          string
	ChainID       uint64
	ChainAddress  string
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
		arg.ChainAddress,
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
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id, chain_address FROM tokens
ORDER BY id ASC 
LIMIT ?
`

func (q *Queries) ListTokens(ctx context.Context, limit int32) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, listTokens, limit)
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
			&i.ChainAddress,
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

const nextTokensPage = `-- name: NextTokensPage :many
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id, chain_address FROM tokens
WHERE id >= ?
ORDER BY id ASC 
LIMIT ?
`

type NextTokensPageParams struct {
	PageCursor annotations.Address
	Limit      int32
}

func (q *Queries) NextTokensPage(ctx context.Context, arg NextTokensPageParams) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, nextTokensPage, arg.PageCursor, arg.Limit)
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
			&i.ChainAddress,
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

const prevTokensPage = `-- name: PrevTokensPage :many
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id, chain_address FROM (
    SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id, chain_address FROM tokens
    WHERE id <= ?
    ORDER BY id DESC
    LIMIT ?
) as token ORDER BY token.id ASC
`

type PrevTokensPageParams struct {
	PageCursor annotations.Address
	Limit      int32
}

func (q *Queries) PrevTokensPage(ctx context.Context, arg PrevTokensPageParams) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, prevTokensPage, arg.PageCursor, arg.Limit)
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
			&i.ChainAddress,
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
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id, chain_address FROM tokens
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
		&i.ChainAddress,
	)
	return i, err
}

const tokenByIDAndChainID = `-- name: TokenByIDAndChainID :one
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id, chain_address FROM tokens
WHERE id = ? AND chain_id = ?
LIMIT 1
`

type TokenByIDAndChainIDParams struct {
	ID      annotations.Address
	ChainID uint64
}

func (q *Queries) TokenByIDAndChainID(ctx context.Context, arg TokenByIDAndChainIDParams) (Token, error) {
	row := q.db.QueryRowContext(ctx, tokenByIDAndChainID, arg.ID, arg.ChainID)
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
		&i.ChainAddress,
	)
	return i, err
}

const tokensByStrategyID = `-- name: TokensByStrategyID :many
SELECT t.id, t.name, t.symbol, t.decimals, t.total_supply, t.creation_block, t.type_id, t.synced, t.tags, t.chain_id, t.chain_address, st.strategy_id, st.token_id, st.min_balance, st.chain_id FROM tokens t
JOIN strategy_tokens st ON st.token_id = t.id
WHERE st.strategy_id = ?
ORDER BY t.name
`

type TokensByStrategyIDRow struct {
	ID            annotations.Address
	Name          string
	Symbol        string
	Decimals      uint64
	TotalSupply   annotations.BigInt
	CreationBlock int64
	TypeID        uint64
	Synced        bool
	Tags          string
	ChainID       uint64
	ChainAddress  string
	StrategyID    uint64
	TokenID       []byte
	MinBalance    []byte
	ChainID_2     uint64
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
			&i.ChainAddress,
			&i.StrategyID,
			&i.TokenID,
			&i.MinBalance,
			&i.ChainID_2,
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
	CreationBlock int64
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
