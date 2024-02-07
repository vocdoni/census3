// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: strategies.sql

package queries

import (
	"context"
	"database/sql"
)

const createStategy = `-- name: CreateStategy :execresult
INSERT INTO strategies (alias, predicate, uri)
VALUES (?, ?, ?)
`

type CreateStategyParams struct {
	Alias     string
	Predicate string
	Uri       string
}

func (q *Queries) CreateStategy(ctx context.Context, arg CreateStategyParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createStategy, arg.Alias, arg.Predicate, arg.Uri)
}

const createStrategyToken = `-- name: CreateStrategyToken :execresult
INSERT INTO strategy_tokens (
    strategy_id,
    token_id,
    chain_id,
    min_balance,
    external_id
)
VALUES (
    ?, ?, ?, ?, ?
)
`

type CreateStrategyTokenParams struct {
	StrategyID uint64
	TokenID    []byte
	ChainID    uint64
	MinBalance string
	ExternalID string
}

func (q *Queries) CreateStrategyToken(ctx context.Context, arg CreateStrategyTokenParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createStrategyToken,
		arg.StrategyID,
		arg.TokenID,
		arg.ChainID,
		arg.MinBalance,
		arg.ExternalID,
	)
}

const deleteStrategiesByToken = `-- name: DeleteStrategiesByToken :execresult
DELETE FROM strategies WHERE id IN (
    SELECT strategy_id FROM strategy_tokens WHERE token_id = ? AND chain_id = ? AND external_id = ?
)
`

type DeleteStrategiesByTokenParams struct {
	TokenID    []byte
	ChainID    uint64
	ExternalID string
}

func (q *Queries) DeleteStrategiesByToken(ctx context.Context, arg DeleteStrategiesByTokenParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteStrategiesByToken, arg.TokenID, arg.ChainID, arg.ExternalID)
}

const deleteStrategyTokensByToken = `-- name: DeleteStrategyTokensByToken :execresult
DELETE FROM strategy_tokens WHERE token_id = ? AND chain_id = ? AND external_id = ?
`

type DeleteStrategyTokensByTokenParams struct {
	TokenID    []byte
	ChainID    uint64
	ExternalID string
}

func (q *Queries) DeleteStrategyTokensByToken(ctx context.Context, arg DeleteStrategyTokensByTokenParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteStrategyTokensByToken, arg.TokenID, arg.ChainID, arg.ExternalID)
}

const existsStrategyByURI = `-- name: ExistsStrategyByURI :one
SELECT EXISTS(SELECT 1 FROM strategies WHERE uri = ?)
`

func (q *Queries) ExistsStrategyByURI(ctx context.Context, uri string) (bool, error) {
	row := q.db.QueryRowContext(ctx, existsStrategyByURI, uri)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const listStrategies = `-- name: ListStrategies :many
SELECT id, predicate, alias, uri FROM strategies
ORDER BY id
`

func (q *Queries) ListStrategies(ctx context.Context) ([]Strategy, error) {
	rows, err := q.db.QueryContext(ctx, listStrategies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Strategy
	for rows.Next() {
		var i Strategy
		if err := rows.Scan(
			&i.ID,
			&i.Predicate,
			&i.Alias,
			&i.Uri,
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

const nextStrategiesPage = `-- name: NextStrategiesPage :many
SELECT id, predicate, alias, uri FROM strategies
WHERE id >= ?
ORDER BY id ASC 
LIMIT ?
`

type NextStrategiesPageParams struct {
	PageCursor uint64
	Limit      int32
}

func (q *Queries) NextStrategiesPage(ctx context.Context, arg NextStrategiesPageParams) ([]Strategy, error) {
	rows, err := q.db.QueryContext(ctx, nextStrategiesPage, arg.PageCursor, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Strategy
	for rows.Next() {
		var i Strategy
		if err := rows.Scan(
			&i.ID,
			&i.Predicate,
			&i.Alias,
			&i.Uri,
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

const prevStrategiesPage = `-- name: PrevStrategiesPage :many
SELECT id, predicate, alias, uri FROM (
    SELECT id, predicate, alias, uri FROM strategies
    WHERE id <= ?
    ORDER BY id DESC 
    LIMIT ?
) as strategy ORDER BY strategy.id ASC
`

type PrevStrategiesPageParams struct {
	PageCursor uint64
	Limit      int32
}

func (q *Queries) PrevStrategiesPage(ctx context.Context, arg PrevStrategiesPageParams) ([]Strategy, error) {
	rows, err := q.db.QueryContext(ctx, prevStrategiesPage, arg.PageCursor, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Strategy
	for rows.Next() {
		var i Strategy
		if err := rows.Scan(
			&i.ID,
			&i.Predicate,
			&i.Alias,
			&i.Uri,
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

const strategiesByTokenIDAndChainIDAndExternalID = `-- name: StrategiesByTokenIDAndChainIDAndExternalID :many
SELECT s.id, s.predicate, s.alias, s.uri FROM strategies s
JOIN strategy_tokens st ON st.strategy_id = s.id
WHERE st.token_id = ? AND st.chain_id = ? AND st.external_id = ?
ORDER BY s.id
`

type StrategiesByTokenIDAndChainIDAndExternalIDParams struct {
	TokenID    []byte
	ChainID    uint64
	ExternalID string
}

func (q *Queries) StrategiesByTokenIDAndChainIDAndExternalID(ctx context.Context, arg StrategiesByTokenIDAndChainIDAndExternalIDParams) ([]Strategy, error) {
	rows, err := q.db.QueryContext(ctx, strategiesByTokenIDAndChainIDAndExternalID, arg.TokenID, arg.ChainID, arg.ExternalID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Strategy
	for rows.Next() {
		var i Strategy
		if err := rows.Scan(
			&i.ID,
			&i.Predicate,
			&i.Alias,
			&i.Uri,
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

const strategyByID = `-- name: StrategyByID :one
SELECT id, predicate, alias, uri FROM strategies
WHERE id = ?
LIMIT 1
`

func (q *Queries) StrategyByID(ctx context.Context, id uint64) (Strategy, error) {
	row := q.db.QueryRowContext(ctx, strategyByID, id)
	var i Strategy
	err := row.Scan(
		&i.ID,
		&i.Predicate,
		&i.Alias,
		&i.Uri,
	)
	return i, err
}

const strategyTokens = `-- name: StrategyTokens :many
SELECT st.token_id, st.min_balance, st.chain_id, st.external_id, t.chain_address, t.symbol
FROM strategy_tokens st
JOIN tokens t ON st.token_id = t.id AND st.chain_id = t.chain_id AND st.external_id = t.external_id
WHERE st.strategy_id = ?
`

type StrategyTokensRow struct {
	TokenID      []byte
	MinBalance   string
	ChainID      uint64
	ExternalID   string
	ChainAddress string
	Symbol       string
}

func (q *Queries) StrategyTokens(ctx context.Context, strategyID uint64) ([]StrategyTokensRow, error) {
	rows, err := q.db.QueryContext(ctx, strategyTokens, strategyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StrategyTokensRow
	for rows.Next() {
		var i StrategyTokensRow
		if err := rows.Scan(
			&i.TokenID,
			&i.MinBalance,
			&i.ChainID,
			&i.ExternalID,
			&i.ChainAddress,
			&i.Symbol,
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

const strategyTokensByStrategyID = `-- name: StrategyTokensByStrategyID :many
SELECT st.token_id as id, st.min_balance, t.symbol, t.chain_address, t.chain_id, t.external_id
FROM strategy_tokens st
JOIN tokens t ON t.id = st.token_id AND t.chain_id = st.chain_id AND t.external_id = st.external_id
WHERE st.strategy_id = ?
ORDER BY strategy_id, token_id
`

type StrategyTokensByStrategyIDRow struct {
	ID           []byte
	MinBalance   string
	Symbol       string
	ChainAddress string
	ChainID      uint64
	ExternalID   string
}

func (q *Queries) StrategyTokensByStrategyID(ctx context.Context, strategyID uint64) ([]StrategyTokensByStrategyIDRow, error) {
	rows, err := q.db.QueryContext(ctx, strategyTokensByStrategyID, strategyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StrategyTokensByStrategyIDRow
	for rows.Next() {
		var i StrategyTokensByStrategyIDRow
		if err := rows.Scan(
			&i.ID,
			&i.MinBalance,
			&i.Symbol,
			&i.ChainAddress,
			&i.ChainID,
			&i.ExternalID,
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

const updateStrategyIPFSUri = `-- name: UpdateStrategyIPFSUri :execresult
UPDATE strategies SET uri = ? WHERE id = ?
`

type UpdateStrategyIPFSUriParams struct {
	Uri string
	ID  uint64
}

func (q *Queries) UpdateStrategyIPFSUri(ctx context.Context, arg UpdateStrategyIPFSUriParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateStrategyIPFSUri, arg.Uri, arg.ID)
}
