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

-- name: TokenByIDAndChainID :one
SELECT * FROM tokens
WHERE id = ? AND chain_id = ?
LIMIT 1;

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
    chain_id,
    chain_address
)
VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: UpdateTokenStatus :execresult
UPDATE tokens
SET synced = sqlc.arg(synced)
WHERE id = sqlc.arg(id);

-- name: UpdateTokenCreationBlock :execresult
UPDATE tokens
SET creation_block = sqlc.arg(creation_block)
WHERE id = sqlc.arg(id);

-- name: ExistsToken :one
SELECT EXISTS 
    (SELECT id 
    FROM tokens
    WHERE id = ?);

-- name: ExistsTokenByChainID :one
SELECT EXISTS 
    (SELECT id 
    FROM tokens
    WHERE id = ? AND chain_id = ?);
