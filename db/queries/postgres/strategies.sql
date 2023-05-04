-- name: PaginatedStrategies :many
SELECT * FROM Strategies
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: StrategyByID :one
SELECT * FROM Strategies
WHERE id = $1
LIMIT 1;

-- name: StrategyByPredicate :one
SELECT * FROM Strategies
WHERE predicate = $1
LIMIT 1;

-- name: PaginatedStrategiesByTokenID :many
SELECT s.* FROM Strategies s
JOIN StrategyTokens st ON st.strategy_id = s.id
WHERE st.token_id = $1
ORDER BY s.id
LIMIT $2 OFFSET $3;

-- name: CreateStategy :execresult
INSERT INTO Strategies (predicate)
VALUES ($1);

-- name: UpdateStrategy :execresult
UPDATE Strategies
SET predicate = sqlc.arg(predicate)
WHERE id = sqlc.arg(id);

-- name: DeleteStrategy :execresult
DELETE FROM Strategies
WHERE id = $1;

-- name: CreateStrategyToken :execresult
INSERT INTO StrategyTokens (
    strategy_id,
    token_id,
    min_balance,
    method_hash
)
VALUES (
    $1, $2, $3, $4
);

-- name: UpdateStrategyToken :execresult
UPDATE StrategyTokens
SET min_balance = sqlc.arg(min_balance),
    method_hash = sqlc.arg(method_hash)
WHERE strategy_id = sqlc.arg(strategy_id) AND token_id = sqlc.arg(token_id);

-- name: DeleteStrategyToken :execresult
DELETE FROM StrategyTokens
WHERE strategy_id = $1 AND token_id = $2;

-- name: PaginatedStrategyTokens :many
SELECT *
FROM StrategyTokens
ORDER BY strategy_id, token_id
LIMIT $1 OFFSET $2;

-- name: StrategyTokenByStrategyIDAndTokenID :one 
SELECT *
FROM StrategyTokens
WHERE strategy_id = $1 AND token_id = $2
LIMIT 1;

-- name: StrategyTokenByStrategyIDAndTokenIDAndMethodHash :one
SELECT *
FROM StrategyTokens
WHERE strategy_id = $1 AND token_id = $2 AND method_hash = $3;
