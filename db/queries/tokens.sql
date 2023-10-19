-- name: ListTokens :many
SELECT * FROM tokens
ORDER BY id ASC 
LIMIT ?;

-- name: NextTokensPage :many
SELECT * FROM tokens
WHERE id >= sqlc.arg(page_cursor)
ORDER BY id ASC 
LIMIT ?;

-- name: PrevTokensPage :many
SELECT * FROM (
    SELECT * FROM tokens
    WHERE id <= sqlc.arg(page_cursor)
    ORDER BY id DESC
    LIMIT ?
) as token ORDER BY token.id ASC;

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
    type_id,
    synced,
    tags,
    chain_id
)
VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: UpdateToken :execresult
UPDATE tokens
SET name = sqlc.arg(name),
    symbol = sqlc.arg(symbol),
    decimals = sqlc.arg(decimals),
    total_supply = sqlc.arg(total_supply),
    creation_block = sqlc.arg(creation_block),
    type_id = sqlc.arg(type_id),
    synced = sqlc.arg(synced),
    tags = sqlc.arg(tags)
WHERE id = sqlc.arg(id);

-- name: UpdateTokenStatus :execresult
UPDATE tokens
SET synced = sqlc.arg(synced)
WHERE id = sqlc.arg(id);

-- name: UpdateTokenCreationBlock :execresult
UPDATE tokens
SET creation_block = sqlc.arg(creation_block)
WHERE id = sqlc.arg(id);

-- name: DeleteToken :execresult
DELETE FROM tokens
WHERE id = ?;

-- name: ExistsToken :one
SELECT EXISTS 
    (SELECT id 
    FROM tokens
    WHERE id = ?)
