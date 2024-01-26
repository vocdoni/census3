-- name: ListLastNoSyncedTokens :many
SELECT * FROM tokens 
WHERE strftime('%s', 'now') - strftime('%s', created_at) <= 600
    AND synced = 0
ORDER BY created_at DESC;

-- name: ListOldNoSyncedTokens :many
SELECT * FROM tokens 
WHERE strftime('%s', 'now') - strftime('%s', created_at) > 600
    AND synced = 0;

-- name: ListSyncedTokens :many
SELECT * FROM tokens WHERE synced = 1;

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

-- name: GetToken :one
SELECT * FROM tokens
WHERE id = ? AND chain_id = ? AND external_id = ?
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
    chain_address,
    external_id,
    default_strategy,
    icon_uri
)
VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 0, ?
);

-- name: UpdateTokenStatus :execresult
UPDATE tokens
SET synced = sqlc.arg(synced), 
    last_block = sqlc.arg(last_block),
    analysed_transfers = sqlc.arg(analysed_transfers)
WHERE id = sqlc.arg(id) 
    AND chain_id = sqlc.arg(chain_id) 
    AND external_id = sqlc.arg(external_id);


-- name: UpdateTokenBlocks :execresult
UPDATE tokens
SET creation_block = sqlc.arg(creation_block),
    last_block = sqlc.arg(last_block)
WHERE id = sqlc.arg(id)
    AND chain_id = sqlc.arg(chain_id)
    AND external_id = sqlc.arg(external_id);

-- name: UpdateTokenDefaultStrategy :execresult
UPDATE tokens
SET default_strategy = sqlc.arg(default_strategy)
WHERE id = sqlc.arg(id)
    AND chain_id = sqlc.arg(chain_id)
    AND external_id = sqlc.arg(external_id);

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

-- name: ExistsTokenByChainIDAndExternalID :one
SELECT EXISTS 
    (SELECT id 
    FROM tokens
    WHERE id = ? AND chain_id = ? AND external_id = ?);

-- name: ExistsAndUnique :one
SELECT COUNT(*) AS num_of_tokens
FROM tokens WHERE id = ? AND chain_id = ? AND external_id = ?
HAVING num_of_tokens = 1;

-- name: DeleteToken :execresult
DELETE FROM tokens WHERE id = ? AND chain_id = ? AND external_id = ?;