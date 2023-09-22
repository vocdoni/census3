-- name: ListStrategies :many
SELECT * FROM strategies
ORDER BY id;

-- name: StrategyByID :one
SELECT * FROM strategies
WHERE id = ?
LIMIT 1;

-- name: StrategiesByTokenID :many
SELECT s.* FROM strategies s
JOIN strategy_tokens st ON st.strategy_id = s.id
WHERE st.token_id = ?
ORDER BY s.id;

-- name: CreateStategy :execresult
INSERT INTO strategies (alias, predicate)
VALUES (?, ?);

-- name: CreateStrategyToken :execresult
INSERT INTO strategy_tokens (
    strategy_id,
    token_id,
    chain_id,
    min_balance
)
VALUES (
    ?, ?, ?, ?
);

-- name: StrategyTokens :many
SELECT *
FROM strategy_tokens
ORDER BY strategy_id, token_id;

-- name: StrategyTokensByStrategyID :many
SELECT st.*, t.symbol
FROM strategy_tokens st
JOIN tokens t ON t.ID = st.token_id
WHERE strategy_id = ?
ORDER BY strategy_id, token_id;