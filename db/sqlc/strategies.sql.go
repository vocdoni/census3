// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: strategies.sql

package queries

import (
	"context"
	"database/sql"
)

const createStategy = `-- name: CreateStategy :execresult
INSERT INTO strategies (alias, predicate)
VALUES (?, ?)
`

type CreateStategyParams struct {
	Alias     string
	Predicate string
}

func (q *Queries) CreateStategy(ctx context.Context, arg CreateStategyParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createStategy, arg.Alias, arg.Predicate)
}

const createStrategyToken = `-- name: CreateStrategyToken :execresult
INSERT INTO strategy_tokens (
    strategy_id,
    token_id,
    chain_id,
    min_balance
)
VALUES (
    ?, ?, ?, ?
)
`

type CreateStrategyTokenParams struct {
	StrategyID uint64
	TokenID    []byte
	ChainID    uint64
	MinBalance []byte
}

func (q *Queries) CreateStrategyToken(ctx context.Context, arg CreateStrategyTokenParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createStrategyToken,
		arg.StrategyID,
		arg.TokenID,
		arg.ChainID,
		arg.MinBalance,
	)
}

const listStrategies = `-- name: ListStrategies :many
SELECT id, predicate, alias FROM strategies
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
		if err := rows.Scan(&i.ID, &i.Predicate, &i.Alias); err != nil {
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

const strategiesByTokenID = `-- name: StrategiesByTokenID :many
SELECT s.id, s.predicate, s.alias FROM strategies s
JOIN strategy_tokens st ON st.strategy_id = s.id
WHERE st.token_id = ?
ORDER BY s.id
`

func (q *Queries) StrategiesByTokenID(ctx context.Context, tokenID []byte) ([]Strategy, error) {
	rows, err := q.db.QueryContext(ctx, strategiesByTokenID, tokenID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Strategy
	for rows.Next() {
		var i Strategy
		if err := rows.Scan(&i.ID, &i.Predicate, &i.Alias); err != nil {
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
SELECT id, predicate, alias FROM strategies
WHERE id = ?
LIMIT 1
`

func (q *Queries) StrategyByID(ctx context.Context, id uint64) (Strategy, error) {
	row := q.db.QueryRowContext(ctx, strategyByID, id)
	var i Strategy
	err := row.Scan(&i.ID, &i.Predicate, &i.Alias)
	return i, err
}

const strategyTokens = `-- name: StrategyTokens :many
SELECT strategy_id, token_id, min_balance, chain_id
FROM strategy_tokens
ORDER BY strategy_id, token_id
`

func (q *Queries) StrategyTokens(ctx context.Context) ([]StrategyToken, error) {
	rows, err := q.db.QueryContext(ctx, strategyTokens)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StrategyToken
	for rows.Next() {
		var i StrategyToken
		if err := rows.Scan(
			&i.StrategyID,
			&i.TokenID,
			&i.MinBalance,
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

const strategyTokensByStrategyID = `-- name: StrategyTokensByStrategyID :many
SELECT st.strategy_id, st.token_id, st.min_balance, st.chain_id, t.symbol
FROM strategy_tokens st
JOIN tokens t ON t.ID = st.token_id
WHERE strategy_id = ?
ORDER BY strategy_id, token_id
`

type StrategyTokensByStrategyIDRow struct {
	StrategyID uint64
	TokenID    []byte
	MinBalance []byte
	ChainID    uint64
	Symbol     string
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
			&i.StrategyID,
			&i.TokenID,
			&i.MinBalance,
			&i.ChainID,
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
