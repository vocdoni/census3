-- name: ListTokens :many
SELECT * FROM tokens
ORDER BY type_id, name;

-- name: ListReadyTokens :many
SELECT * FROM tokens
WHERE creation_block IS NOT NULL
ORDER BY type_id, name;

-- name: ListNotReadyTokens :many
SELECT * FROM tokens
WHERE creation_block IS NULL
ORDER BY type_id, name;

-- name: TokenByID :one
SELECT * FROM tokens
WHERE id = ?
LIMIT 1;

-- name: TokenByName :one
SELECT * FROM tokens
WHERE name = ?
LIMIT 1;

-- name: TokenBySymbol :one
SELECT * FROM tokens
WHERE symbol = ?
LIMIT 1;

-- name: TokensByType :many
SELECT * FROM tokens
WHERE type_id = ?
ORDER BY name;

-- name: TokensByStrategyID :many
SELECT t.*, st.* FROM tokens t
JOIN strategy_tokens st ON st.token_id = t.id
WHERE st.strategy_id = ?
ORDER BY t.name;

-- name: CreateToken :execresult
INSERT INTO tokens (
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
UPDATE tokens
SET name = sqlc.arg(name),
    symbol = sqlc.arg(symbol),
    decimals = sqlc.arg(decimals),
    total_supply = sqlc.arg(total_supply),
    creation_block = sqlc.arg(creation_block),
    type_id = sqlc.arg(type_id)
WHERE id = sqlc.arg(id);

-- name: DeleteToken :execresult
DELETE FROM tokens
WHERE id = ?;