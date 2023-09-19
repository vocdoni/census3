-- name: ListStrategies :many
SELECT * FROM strategies
ORDER BY id;

-- name: StrategyByID :one
SELECT * FROM strategies
WHERE id = ?
LIMIT 1;

-- name: StrategyByPredicate :one
SELECT * FROM strategies
WHERE predicate = ?
LIMIT 1;

-- name: StrategiesByTokenID :many
SELECT s.* FROM strategies s
JOIN strategy_tokens st ON st.strategy_id = s.id
WHERE st.token_id = ?
ORDER BY s.id;

-- name: CreateStategy :execresult
INSERT INTO strategies (predicate)
VALUES (?);

-- name: UpdateStrategy :execresult
UPDATE strategies
SET predicate = sqlc.arg(predicate)
WHERE id = sqlc.arg(id);

-- name: DeleteStrategy :execresult
DELETE FROM strategies
WHERE id = ?;

-- name: CreateStrategyToken :execresult
INSERT INTO strategy_tokens (
    strategy_id,
    token_id,
    min_balance
)
VALUES (
    ?, ?, ?
);

-- name: UpdateStrategyToken :execresult
UPDATE strategy_tokens
SET min_balance = sqlc.arg(min_balance)
WHERE strategy_id = sqlc.arg(strategy_id) AND token_id = sqlc.arg(token_id);

-- name: DeleteStrategyToken :execresult
DELETE FROM strategy_tokens
WHERE strategy_id = ? AND token_id = ?;

-- name: StrategyTokens :many
SELECT *
FROM strategy_tokens
ORDER BY strategy_id, token_id;

-- name: StrategyTokenByStrategyIDAndTokenID :one 
SELECT *
FROM strategy_tokens
WHERE strategy_id = ? AND token_id = ?
LIMIT 1;

-- name: StrategyTokenByStrategyIDAndTokenIDAndMethodHash :one
SELECT *
FROM strategy_tokens
WHERE strategy_id = ? AND token_id = ?;