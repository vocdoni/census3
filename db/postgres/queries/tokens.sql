-- name: ListTokens :many
SELECT * FROM tokens
ORDER BY type_id, name;

-- name: TokenByID :one
SELECT * FROM tokens
WHERE id = $1
LIMIT 1;

-- name: TokenByName :one
SELECT * FROM tokens
WHERE name = $1
LIMIT 1;

-- name: TokenBySymbol :one
SELECT * FROM tokens
WHERE symbol = $1
LIMIT 1;

-- name: TokensByType :many
SELECT * FROM tokens
WHERE type_id = $1
ORDER BY name;

-- name: TokensByStrategyID :many
SELECT t.*, st.* FROM tokens t
JOIN strategy_tokens st ON st.token_id = t.id
WHERE st.strategy_id = $1
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
    $1, $2, $3, $4, $5, $6, $7
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
WHERE id = $1;
