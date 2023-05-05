-- name: ListStrategies :many
SELECT * FROM strategies
ORDER BY id;

-- name: StrategyByID :one
SELECT * FROM strategies
WHERE id = $1
LIMIT 1;

-- name: StrategyByPredicate :one
SELECT * FROM strategies
WHERE predicate = $1
LIMIT 1;

-- name: StrategiesByTokenID :many
SELECT s.* FROM strategies s
JOIN strategy_tokens st ON st.strategy_id = s.id
WHERE st.token_id = $1
ORDER BY s.id;

-- name: CreateStategy :execresult
INSERT INTO strategies (predicate)
VALUES ($1);

-- name: UpdateStrategy :execresult
UPDATE strategies
SET predicate = sqlc.arg(predicate)
WHERE id = sqlc.arg(id);

-- name: DeleteStrategy :execresult
DELETE FROM strategies
WHERE id = $1;

-- name: CreateStrategyToken :execresult
INSERT INTO strategy_tokens (
    strategy_id,
    token_id,
    min_balance,
    method_hash
)
VALUES (
    $1, $2, $3, $4
);

-- name: UpdateStrategyToken :execresult
UPDATE strategy_tokens
SET min_balance = sqlc.arg(min_balance),
    method_hash = sqlc.arg(method_hash)
WHERE strategy_id = sqlc.arg(strategy_id) AND token_id = sqlc.arg(token_id);

-- name: DeleteStrategyToken :execresult
DELETE FROM strategy_tokens
WHERE strategy_id = $1 AND token_id = $2;

-- name: StrategyTokens :many
SELECT *
FROM strategy_tokens
ORDER BY strategy_id, token_id;

-- name: StrategyTokenByStrategyIDAndTokenID :one 
SELECT *
FROM strategy_tokens
WHERE strategy_id = $1 AND token_id = $2
LIMIT 1;

-- name: StrategyTokenByStrategyIDAndTokenIDAndMethodHash :one
SELECT *
FROM strategy_tokens
WHERE strategy_id = $1 AND token_id = $2 AND method_hash = $3;
