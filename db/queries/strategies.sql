-- name: PaginatedStrategies :many
SELECT * FROM Strategies
ORDER BY id
LIMIT ? OFFSET ?;

-- name: StrategyByID :one
SELECT * FROM Strategies
WHERE id = ?
LIMIT 1;

-- name: StrategyByPredicate :one
SELECT * FROM Strategies
WHERE predicate = ?
LIMIT 1;

-- name: PaginatedStrategiesByTokenID :many
SELECT s.* FROM Strategies s
JOIN StrategyTokens st ON st.strategy_id = s.id
WHERE st.token_id = ?
ORDER BY s.id
LIMIT ? OFFSET ?;

-- name: CreateStategy :execresult
INSERT INTO Strategies (predicate)
VALUES (?);

-- name: UpdateStrategy :execresult
UPDATE Strategies
SET predicate = sqlc.arg(predicate)
WHERE id = sqlc.arg(id);

-- name: DeleteStrategy :execresult
DELETE FROM Strategies
WHERE id = ?;

-- name: CreateStrategyToken :execresult
INSERT INTO StrategyTokens (
    strategy_id,
    token_id,
    min_balance,
    method_hash
)
VALUES (
    ?, ?, ?, ?
);

-- name: UpdateStrategyToken :execresult
UPDATE StrategyTokens
SET min_balance = sqlc.arg(min_balance),
    method_hash = sqlc.arg(method_hash)
WHERE strategy_id = sqlc.arg(strategy_id) AND token_id = sqlc.arg(token_id);

-- name: DeleteStrategyToken :execresult
DELETE FROM StrategyTokens
WHERE strategy_id = ? AND token_id = ?;

-- name: PaginatedStrategyTokens :many
SELECT *
FROM StrategyTokens
ORDER BY strategy_id, token_id
LIMIT ? OFFSET ?;

-- name: StrategyTokenByStrategyIDAndTokenID :one 
SELECT *
FROM StrategyTokens
WHERE strategy_id = ? AND token_id = ?
LIMIT 1;

-- name: StrategyTokenByStrategyIDAndTokenIDAndMethodHash :one
SELECT *
FROM StrategyTokens
WHERE strategy_id = ? AND token_id = ? AND method_hash = ?;
