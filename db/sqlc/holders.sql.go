// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: holders.sql

package queries

import (
	"context"
	"database/sql"

	"github.com/vocdoni/census3/db/annotations"
)

const aNDOperator = `-- name: ANDOperator :many
;WITH holders_a as (
    SELECT th.holder_id, th.balance
    FROM token_holders th
    WHERE th.token_id = ? 
        AND th.chain_id = ?
        AND th.external_id = ?
        AND th.balance >= ?
),
holders_b as (
    SELECT th.holder_id, th.balance
    FROM token_holders th
    WHERE th.token_id = ? 
        AND th.chain_id = ?
        AND th.external_id = ?
        AND th.balance >= ?
)
SELECT holders_a.holder_id, IFNULL(holders_a.balance, '0') as balance_a, IFNULL(holders_b.balance, '0') as balance_b
FROM holders_a
INNER JOIN holders_b ON holders_a.holder_id = holders_b.holder_id
`

type ANDOperatorParams struct {
	TokenIDA    annotations.Address
	ChainIDA    uint64
	ExternalIDA string
	MinBalanceA string
	TokenIDB    annotations.Address
	ChainIDB    uint64
	ExternalIDB string
	MinBalanceB string
}

type ANDOperatorRow struct {
	HolderID []byte
	BalanceA interface{}
	BalanceB interface{}
}

func (q *Queries) ANDOperator(ctx context.Context, arg ANDOperatorParams) ([]ANDOperatorRow, error) {
	rows, err := q.db.QueryContext(ctx, aNDOperator,
		arg.TokenIDA,
		arg.ChainIDA,
		arg.ExternalIDA,
		arg.MinBalanceA,
		arg.TokenIDB,
		arg.ChainIDB,
		arg.ExternalIDB,
		arg.MinBalanceB,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ANDOperatorRow
	for rows.Next() {
		var i ANDOperatorRow
		if err := rows.Scan(&i.HolderID, &i.BalanceA, &i.BalanceB); err != nil {
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

const countTokenHolders = `-- name: CountTokenHolders :one
SELECT COUNT(holder_id) 
FROM token_holders
WHERE token_id = ?
    AND chain_id = ?
    AND external_id = ?
    AND balance >= ?
`

type CountTokenHoldersParams struct {
	TokenID    annotations.Address
	ChainID    uint64
	ExternalID string
	Balance    string
}

func (q *Queries) CountTokenHolders(ctx context.Context, arg CountTokenHoldersParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countTokenHolders,
		arg.TokenID,
		arg.ChainID,
		arg.ExternalID,
		arg.Balance,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createTokenHolder = `-- name: CreateTokenHolder :execresult
INSERT INTO token_holders (
    token_id,
    holder_id,
    balance,
    block_id,
    chain_id,
    external_id
)
VALUES (
    ?, ?, ?, ?, ?, ?
)
`

type CreateTokenHolderParams struct {
	TokenID    annotations.Address
	HolderID   annotations.Address
	Balance    string
	BlockID    uint64
	ChainID    uint64
	ExternalID string
}

func (q *Queries) CreateTokenHolder(ctx context.Context, arg CreateTokenHolderParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createTokenHolder,
		arg.TokenID,
		arg.HolderID,
		arg.Balance,
		arg.BlockID,
		arg.ChainID,
		arg.ExternalID,
	)
}

const deleteTokenHolder = `-- name: DeleteTokenHolder :execresult
DELETE FROM token_holders WHERE token_id = ? AND chain_id = ? AND external_id = ?
`

type DeleteTokenHolderParams struct {
	TokenID    annotations.Address
	ChainID    uint64
	ExternalID string
}

func (q *Queries) DeleteTokenHolder(ctx context.Context, arg DeleteTokenHolderParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteTokenHolder, arg.TokenID, arg.ChainID, arg.ExternalID)
}

const existTokenHolder = `-- name: ExistTokenHolder :one
SELECT EXISTS (
    SELECT holder_id 
    FROM token_holders
    WHERE token_id = ? 
        AND holder_id = ?
        AND chain_id = ?
        AND external_id = ?
)
`

type ExistTokenHolderParams struct {
	TokenID    annotations.Address
	HolderID   annotations.Address
	ChainID    uint64
	ExternalID string
}

func (q *Queries) ExistTokenHolder(ctx context.Context, arg ExistTokenHolderParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, existTokenHolder,
		arg.TokenID,
		arg.HolderID,
		arg.ChainID,
		arg.ExternalID,
	)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getTokenHolder = `-- name: GetTokenHolder :one
SELECT token_id, holder_id, balance, block_id, chain_id, external_id
FROM token_holders
WHERE token_id = ? 
    AND holder_id = ? 
    AND chain_id = ?
    AND external_id = ?
`

type GetTokenHolderParams struct {
	TokenID    annotations.Address
	HolderID   annotations.Address
	ChainID    uint64
	ExternalID string
}

func (q *Queries) GetTokenHolder(ctx context.Context, arg GetTokenHolderParams) (TokenHolder, error) {
	row := q.db.QueryRowContext(ctx, getTokenHolder,
		arg.TokenID,
		arg.HolderID,
		arg.ChainID,
		arg.ExternalID,
	)
	var i TokenHolder
	err := row.Scan(
		&i.TokenID,
		&i.HolderID,
		&i.Balance,
		&i.BlockID,
		&i.ChainID,
		&i.ExternalID,
	)
	return i, err
}

const listTokenHolders = `-- name: ListTokenHolders :many
SELECT token_id, holder_id, balance, block_id, chain_id, external_id
FROM token_holders
WHERE token_id = ? AND chain_id = ? AND external_id = ?
`

type ListTokenHoldersParams struct {
	TokenID    annotations.Address
	ChainID    uint64
	ExternalID string
}

func (q *Queries) ListTokenHolders(ctx context.Context, arg ListTokenHoldersParams) ([]TokenHolder, error) {
	rows, err := q.db.QueryContext(ctx, listTokenHolders, arg.TokenID, arg.ChainID, arg.ExternalID)
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

const nextStrategyTokenHoldersPage = `-- name: NextStrategyTokenHoldersPage :many
SELECT token_holders.token_id, token_holders.holder_id, token_holders.balance, token_holders.block_id, token_holders.chain_id, token_holders.external_id
FROM token_holders
JOIN strategy_tokens 
    ON strategy_tokens.token_id = token_holders.token_id
    AND strategy_tokens.chain_id = token_holders.chain_id
    AND strategy_tokens.external_id = token_holders.external_id
WHERE strategy_tokens.strategy_id = ?
    AND strategy_tokens.min_balance <= token_holders.balance
    AND token_holders.holder_id >= ?
ORDER BY token_holders.holder_id ASC 
LIMIT ?
`

type NextStrategyTokenHoldersPageParams struct {
	StrategyID uint64
	PageCursor annotations.Address
	Limit      int32
}

func (q *Queries) NextStrategyTokenHoldersPage(ctx context.Context, arg NextStrategyTokenHoldersPageParams) ([]TokenHolder, error) {
	rows, err := q.db.QueryContext(ctx, nextStrategyTokenHoldersPage, arg.StrategyID, arg.PageCursor, arg.Limit)
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

const oROperator = `-- name: OROperator :many
SELECT holder_ids.holder_id, IFNULL(a.balance, '0') AS balance_a, IFNULL(b.balance, '0') AS balance_b
FROM (
    SELECT th.holder_id
    FROM token_holders th
    WHERE (
        th.token_id = ? 
        AND th.chain_id = ?
        AND th.external_id = ?
        AND th.balance >= ?
    ) OR (
        th.token_id = ? 
        AND th.chain_id = ?
        AND th.external_id = ?
        AND th.balance >= ?
    )
) as holder_ids
LEFT JOIN (
    SELECT th_b.holder_id, th_b.balance
    FROM token_holders th_b
    WHERE th_b.token_id = ? 
        AND th_b.chain_id = ?
        AND th_b.external_id = ?
        AND th_b.balance >= ?
) AS a ON holder_ids.holder_id = a.holder_id
LEFT JOIN (
    SELECT th_a.holder_id, th_a.balance
    FROM token_holders th_a
    WHERE th_a.token_id = ? 
        AND th_a.chain_id = ?
        AND th_a.external_id = ?
        AND th_a.balance >= ?
) AS b ON holder_ids.holder_id = b.holder_id
GROUP BY holder_ids.holder_id
`

type OROperatorParams struct {
	TokenIDA    annotations.Address
	ChainIDA    uint64
	ExternalIDA string
	MinBalanceA string
	TokenIDB    annotations.Address
	ChainIDB    uint64
	ExternalIDB string
	MinBalanceB string
}

type OROperatorRow struct {
	HolderID annotations.Address
	BalanceA interface{}
	BalanceB interface{}
}

func (q *Queries) OROperator(ctx context.Context, arg OROperatorParams) ([]OROperatorRow, error) {
	rows, err := q.db.QueryContext(ctx, oROperator,
		arg.TokenIDA,
		arg.ChainIDA,
		arg.ExternalIDA,
		arg.MinBalanceA,
		arg.TokenIDB,
		arg.ChainIDB,
		arg.ExternalIDB,
		arg.MinBalanceB,
		arg.TokenIDA,
		arg.ChainIDA,
		arg.ExternalIDA,
		arg.MinBalanceA,
		arg.TokenIDB,
		arg.ChainIDB,
		arg.ExternalIDB,
		arg.MinBalanceB,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OROperatorRow
	for rows.Next() {
		var i OROperatorRow
		if err := rows.Scan(&i.HolderID, &i.BalanceA, &i.BalanceB); err != nil {
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

const prevStrategyTokenHoldersPage = `-- name: PrevStrategyTokenHoldersPage :many
SELECT token_id, holder_id, balance, block_id, chain_id, external_id FROM (
    SELECT token_holders.token_id, token_holders.holder_id, token_holders.balance, token_holders.block_id, token_holders.chain_id, token_holders.external_id
    FROM token_holders
    JOIN strategy_tokens 
        ON strategy_tokens.token_id = token_holders.token_id
        AND strategy_tokens.chain_id = token_holders.chain_id
        AND strategy_tokens.external_id = token_holders.external_id
    WHERE strategy_tokens.strategy_id = ?
        AND strategy_tokens.min_balance <= token_holders.balance
        AND token_holders.holder_id <= ?
    ORDER BY token_holders.holder_id DESC 
    LIMIT ?
) as holder ORDER BY holder.holder_id ASC
`

type PrevStrategyTokenHoldersPageParams struct {
	StrategyID uint64
	PageCursor annotations.Address
	Limit      int32
}

func (q *Queries) PrevStrategyTokenHoldersPage(ctx context.Context, arg PrevStrategyTokenHoldersPageParams) ([]TokenHolder, error) {
	rows, err := q.db.QueryContext(ctx, prevStrategyTokenHoldersPage, arg.StrategyID, arg.PageCursor, arg.Limit)
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

const tokenHoldersByStrategyID = `-- name: TokenHoldersByStrategyID :many
SELECT token_holders.holder_id, token_holders.balance, strategy_tokens.min_balance
FROM token_holders
JOIN strategy_tokens 
    ON strategy_tokens.token_id = token_holders.token_id
    AND strategy_tokens.chain_id = token_holders.chain_id
    AND strategy_tokens.external_id = token_holders.external_id
WHERE strategy_tokens.strategy_id = ?
    AND strategy_tokens.min_balance <= token_holders.balance
`

type TokenHoldersByStrategyIDRow struct {
	HolderID   annotations.Address
	Balance    string
	MinBalance string
}

func (q *Queries) TokenHoldersByStrategyID(ctx context.Context, strategyID uint64) ([]TokenHoldersByStrategyIDRow, error) {
	rows, err := q.db.QueryContext(ctx, tokenHoldersByStrategyID, strategyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TokenHoldersByStrategyIDRow
	for rows.Next() {
		var i TokenHoldersByStrategyIDRow
		if err := rows.Scan(&i.HolderID, &i.Balance, &i.MinBalance); err != nil {
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

const tokenHoldersByTokenIDAndChainIDAndMinBalance = `-- name: TokenHoldersByTokenIDAndChainIDAndMinBalance :many
SELECT token_holders.holder_id, token_holders.balance
FROM token_holders
WHERE token_holders.token_id = ? 
    AND token_holders.chain_id = ?
    AND token_holders.external_id = ?
    AND token_holders.balance >= ?
`

type TokenHoldersByTokenIDAndChainIDAndMinBalanceParams struct {
	TokenID    annotations.Address
	ChainID    uint64
	ExternalID string
	Balance    string
}

type TokenHoldersByTokenIDAndChainIDAndMinBalanceRow struct {
	HolderID annotations.Address
	Balance  string
}

func (q *Queries) TokenHoldersByTokenIDAndChainIDAndMinBalance(ctx context.Context, arg TokenHoldersByTokenIDAndChainIDAndMinBalanceParams) ([]TokenHoldersByTokenIDAndChainIDAndMinBalanceRow, error) {
	rows, err := q.db.QueryContext(ctx, tokenHoldersByTokenIDAndChainIDAndMinBalance,
		arg.TokenID,
		arg.ChainID,
		arg.ExternalID,
		arg.Balance,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TokenHoldersByTokenIDAndChainIDAndMinBalanceRow
	for rows.Next() {
		var i TokenHoldersByTokenIDAndChainIDAndMinBalanceRow
		if err := rows.Scan(&i.HolderID, &i.Balance); err != nil {
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
WHERE token_id = ? 
    AND holder_id = ? 
    AND block_id = ? 
    AND chain_id = ? 
    AND external_id = ?
`

type UpdateTokenHolderBalanceParams struct {
	Balance    string
	NewBlockID uint64
	TokenID    annotations.Address
	HolderID   annotations.Address
	BlockID    uint64
	ChainID    uint64
	ExternalID string
}

func (q *Queries) UpdateTokenHolderBalance(ctx context.Context, arg UpdateTokenHolderBalanceParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateTokenHolderBalance,
		arg.Balance,
		arg.NewBlockID,
		arg.TokenID,
		arg.HolderID,
		arg.BlockID,
		arg.ChainID,
		arg.ExternalID,
	)
}
