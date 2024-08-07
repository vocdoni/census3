// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: tokens.sql

package queries

import (
	"context"
	"database/sql"
	"time"

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
    chain_address,
    external_id,
    default_strategy,
    icon_uri,
    last_block
)
VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 0, ?, ?
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
	ExternalID    string
	IconUri       string
	LastBlock     int64
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
		arg.ExternalID,
		arg.IconUri,
		arg.LastBlock,
	)
}

const deleteToken = `-- name: DeleteToken :execresult
DELETE FROM tokens WHERE id = ? AND chain_id = ? AND external_id = ?
`

type DeleteTokenParams struct {
	ID         annotations.Address
	ChainID    uint64
	ExternalID string
}

func (q *Queries) DeleteToken(ctx context.Context, arg DeleteTokenParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteToken, arg.ID, arg.ChainID, arg.ExternalID)
}

const existsAndUnique = `-- name: ExistsAndUnique :one
SELECT COUNT(*) AS num_of_tokens
FROM tokens WHERE id = ? AND chain_id = ? AND external_id = ?
HAVING num_of_tokens = 1
`

type ExistsAndUniqueParams struct {
	ID         annotations.Address
	ChainID    uint64
	ExternalID string
}

func (q *Queries) ExistsAndUnique(ctx context.Context, arg ExistsAndUniqueParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, existsAndUnique, arg.ID, arg.ChainID, arg.ExternalID)
	var num_of_tokens int64
	err := row.Scan(&num_of_tokens)
	return num_of_tokens, err
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

const existsTokenByChainIDAndExternalID = `-- name: ExistsTokenByChainIDAndExternalID :one
SELECT EXISTS 
    (SELECT id 
    FROM tokens
    WHERE id = ? AND chain_id = ? AND external_id = ?)
`

type ExistsTokenByChainIDAndExternalIDParams struct {
	ID         annotations.Address
	ChainID    uint64
	ExternalID string
}

func (q *Queries) ExistsTokenByChainIDAndExternalID(ctx context.Context, arg ExistsTokenByChainIDAndExternalIDParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, existsTokenByChainIDAndExternalID, arg.ID, arg.ChainID, arg.ExternalID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getToken = `-- name: GetToken :one
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id, chain_address, external_id, default_strategy, icon_uri, created_at, last_block, analysed_transfers FROM tokens
WHERE id = ? AND chain_id = ? AND external_id = ?
LIMIT 1
`

type GetTokenParams struct {
	ID         annotations.Address
	ChainID    uint64
	ExternalID string
}

func (q *Queries) GetToken(ctx context.Context, arg GetTokenParams) (Token, error) {
	row := q.db.QueryRowContext(ctx, getToken, arg.ID, arg.ChainID, arg.ExternalID)
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
		&i.ExternalID,
		&i.DefaultStrategy,
		&i.IconUri,
		&i.CreatedAt,
		&i.LastBlock,
		&i.AnalysedTransfers,
	)
	return i, err
}

const listLastNoSyncedTokens = `-- name: ListLastNoSyncedTokens :many
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id, chain_address, external_id, default_strategy, icon_uri, created_at, last_block, analysed_transfers FROM tokens 
WHERE strftime('%s', 'now') - strftime('%s', created_at) <= 600
    AND synced = 0
ORDER BY created_at DESC
`

func (q *Queries) ListLastNoSyncedTokens(ctx context.Context) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, listLastNoSyncedTokens)
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
			&i.ExternalID,
			&i.DefaultStrategy,
			&i.IconUri,
			&i.CreatedAt,
			&i.LastBlock,
			&i.AnalysedTransfers,
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

const listOldNoSyncedTokens = `-- name: ListOldNoSyncedTokens :many
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id, chain_address, external_id, default_strategy, icon_uri, created_at, last_block, analysed_transfers FROM tokens 
WHERE strftime('%s', 'now') - strftime('%s', created_at) > 600
    AND synced = 0
`

func (q *Queries) ListOldNoSyncedTokens(ctx context.Context) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, listOldNoSyncedTokens)
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
			&i.ExternalID,
			&i.DefaultStrategy,
			&i.IconUri,
			&i.CreatedAt,
			&i.LastBlock,
			&i.AnalysedTransfers,
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

const listSyncedTokens = `-- name: ListSyncedTokens :many
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id, chain_address, external_id, default_strategy, icon_uri, created_at, last_block, analysed_transfers FROM tokens WHERE synced = 1
`

func (q *Queries) ListSyncedTokens(ctx context.Context) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, listSyncedTokens)
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
			&i.ExternalID,
			&i.DefaultStrategy,
			&i.IconUri,
			&i.CreatedAt,
			&i.LastBlock,
			&i.AnalysedTransfers,
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
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id, chain_address, external_id, default_strategy, icon_uri, created_at, last_block, analysed_transfers FROM tokens
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
			&i.ExternalID,
			&i.DefaultStrategy,
			&i.IconUri,
			&i.CreatedAt,
			&i.LastBlock,
			&i.AnalysedTransfers,
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
SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id, chain_address, external_id, default_strategy, icon_uri, created_at, last_block, analysed_transfers FROM (
    SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id, chain_address, external_id, default_strategy, icon_uri, created_at, last_block, analysed_transfers FROM tokens
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
			&i.ExternalID,
			&i.DefaultStrategy,
			&i.IconUri,
			&i.CreatedAt,
			&i.LastBlock,
			&i.AnalysedTransfers,
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

const tokensByStrategyID = `-- name: TokensByStrategyID :many
SELECT t.id, t.name, t.symbol, t.decimals, t.total_supply, t.creation_block, t.type_id, t.synced, t.tags, t.chain_id, t.chain_address, t.external_id, t.default_strategy, t.icon_uri, t.created_at, t.last_block, t.analysed_transfers, st.strategy_id, st.token_id, st.min_balance, st.chain_id, st.external_id, st.token_alias FROM tokens t
JOIN strategy_tokens st 
ON st.token_id = t.id 
    AND st.chain_id = t.chain_id 
    AND st.external_id = t.external_id
WHERE st.strategy_id = ?
ORDER BY t.name
`

type TokensByStrategyIDRow struct {
	ID                annotations.Address
	Name              string
	Symbol            string
	Decimals          uint64
	TotalSupply       annotations.BigInt
	CreationBlock     int64
	TypeID            uint64
	Synced            bool
	Tags              string
	ChainID           uint64
	ChainAddress      string
	ExternalID        string
	DefaultStrategy   uint64
	IconUri           string
	CreatedAt         time.Time
	LastBlock         int64
	AnalysedTransfers int64
	StrategyID        uint64
	TokenID           []byte
	MinBalance        string
	ChainID_2         uint64
	ExternalID_2      string
	TokenAlias        string
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
			&i.ExternalID,
			&i.DefaultStrategy,
			&i.IconUri,
			&i.CreatedAt,
			&i.LastBlock,
			&i.AnalysedTransfers,
			&i.StrategyID,
			&i.TokenID,
			&i.MinBalance,
			&i.ChainID_2,
			&i.ExternalID_2,
			&i.TokenAlias,
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

const updateTokenBlocks = `-- name: UpdateTokenBlocks :execresult
UPDATE tokens
SET creation_block = ?,
    last_block = ?
WHERE id = ?
    AND chain_id = ?
    AND external_id = ?
`

type UpdateTokenBlocksParams struct {
	CreationBlock int64
	LastBlock     int64
	ID            annotations.Address
	ChainID       uint64
	ExternalID    string
}

func (q *Queries) UpdateTokenBlocks(ctx context.Context, arg UpdateTokenBlocksParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateTokenBlocks,
		arg.CreationBlock,
		arg.LastBlock,
		arg.ID,
		arg.ChainID,
		arg.ExternalID,
	)
}

const updateTokenDefaultStrategy = `-- name: UpdateTokenDefaultStrategy :execresult
UPDATE tokens
SET default_strategy = ?
WHERE id = ?
    AND chain_id = ?
    AND external_id = ?
`

type UpdateTokenDefaultStrategyParams struct {
	DefaultStrategy uint64
	ID              annotations.Address
	ChainID         uint64
	ExternalID      string
}

func (q *Queries) UpdateTokenDefaultStrategy(ctx context.Context, arg UpdateTokenDefaultStrategyParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateTokenDefaultStrategy,
		arg.DefaultStrategy,
		arg.ID,
		arg.ChainID,
		arg.ExternalID,
	)
}

const updateTokenStatus = `-- name: UpdateTokenStatus :execresult
UPDATE tokens
SET synced = ?, 
    last_block = ?,
    analysed_transfers = ?,
    total_supply = ?
WHERE id = ? 
    AND chain_id = ? 
    AND external_id = ?
`

type UpdateTokenStatusParams struct {
	Synced            bool
	LastBlock         int64
	AnalysedTransfers int64
	TotalSupply       annotations.BigInt
	ID                annotations.Address
	ChainID           uint64
	ExternalID        string
}

func (q *Queries) UpdateTokenStatus(ctx context.Context, arg UpdateTokenStatusParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateTokenStatus,
		arg.Synced,
		arg.LastBlock,
		arg.AnalysedTransfers,
		arg.TotalSupply,
		arg.ID,
		arg.ChainID,
		arg.ExternalID,
	)
}
