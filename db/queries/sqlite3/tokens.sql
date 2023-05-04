-- name: PaginatedTokens :many
SELECT * FROM Tokens
ORDER BY type_id, name
LIMIT ? OFFSET ?;

-- name: TokenByID :one
SELECT * FROM Tokens
WHERE id = ?
LIMIT 1;

-- name: TokenByName :one
SELECT * FROM Tokens
WHERE name = ?
LIMIT 1;

-- name: TokenBySymbol :one
SELECT * FROM Tokens
WHERE symbol = ?
LIMIT 1;

-- name: TokensByType :many
SELECT * FROM Tokens
WHERE type_id = ?
ORDER BY name
LIMIT ? OFFSET ?;

-- name: TokensByStrategyID :many
SELECT t.*, st.* FROM Tokens t
JOIN StrategyTokens st ON st.token_id = t.id
WHERE st.strategy_id = ?
ORDER BY t.name
LIMIT ? OFFSET ?;

-- name: CreateToken :execresult
INSERT INTO Tokens (
    id,
    name,
    symbol,
    decimals,
    total_supply,
    creation_block,
    type_id
)
VALUES (
    ?, ?, ?, ?, ?, ?, ?
);

-- name: UpdateToken :execresult
UPDATE Tokens
SET name = sqlc.arg(name),
    symbol = sqlc.arg(symbol),
    decimals = sqlc.arg(decimals),
    total_supply = sqlc.arg(total_supply),
    creation_block = sqlc.arg(creation_block),
    type_id = sqlc.arg(type_id)
WHERE id = sqlc.arg(id);

-- name: DeleteToken :execresult
DELETE FROM Tokens
WHERE id = ?;
