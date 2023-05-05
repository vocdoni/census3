-- name: ListHolders :many
SELECT * FROM holders
ORDER BY id;

-- name: HolderByID :one
SELECT * FROM holders
WHERE id = $1
LIMIT 1;

-- name: CreateHolder :execresult
INSERT INTO holders (id)
VALUES ($1)
ON CONFLICT DO NOTHING;

-- name: DeleteHolder :execresult
DELETE FROM holders
WHERE id = $1;

-- name: ListTokenHolders :many
SELECT * FROM token_holders
ORDER BY token_id, holder_id, block_id;

-- name: TokensByHolderID :many
SELECT tokens.*
FROM tokens
JOIN token_holders ON Tokens.id = token_holders.token_id
WHERE token_holders.holder_id = $1
LIMIT $2 OFFSET $3;

-- name: TokensByHolderIDAndBlockID :many
SELECT tokens.*
FROM tokens
JOIN token_holders ON tokens.id = token_holders.token_id
WHERE token_holders.holder_id = $1 AND token_holders.block_id = $2
LIMIT $3 OFFSET $4;

-- name: TokenHoldersByTokenID :many
SELECT holders.*, token_holders.balance
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = $1;

-- name: TokenHoldersByTokenIDAndBlockID :many
SELECT holders.*
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = $1 AND token_holders.block_id = $2
LIMIT $3 OFFSET $4;

-- name: TokenHoldersByTokenIDAndMinBalance :many
SELECT holders.*
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = $1 AND token_holders.balance >= $2
LIMIT $3 OFFSET $4;

-- name: TokenHoldersByTokenIDAndBlockIDAndMinBalance :many
SELECT holders.*
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = $1 AND token_holders.balance >= $2 AND token_holders.block_id = $3
LIMIT $4 OFFSET $5;

-- name: TokenHolderByTokenIDAndHolderID :one
SELECT holders.*, token_holders.balance
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = $1 AND token_holders.holder_id = $2;

-- name: TokenHolderByTokenIDAndBlockIDAndHolderID :one
SELECT holders.*, token_holders.balance
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = $1 AND token_holders.holder_id = $2 AND token_holders.block_id = $3;

-- name: LastBlockByTokenID :one
SELECT block_id 
FROM token_holders
WHERE token_id = $1
ORDER BY block_id DESC
LIMIT 1;

-- name: CountTokenHoldersByTokenID :one
SELECT COUNT(holder_id) 
FROM token_holders
WHERE token_id = $1;

-- name: CreateTokenHolder :execresult
INSERT INTO token_holders (
    token_id,
    holder_id,
    balance,
    block_id
)
VALUES (
    $1, $2, $3, $4
);

-- name: UpdateTokenHolder :execresult
UPDATE token_holders
SET balance = sqlc.arg(balance),
    block_id = sqlc.arg(block_id)
WHERE token_id = sqlc.arg(token_id) AND holder_id = sqlc.arg(holder_id) AND block_id = sqlc.arg(block_id);

-- name: DeleteTokenHolder :execresult
DELETE FROM token_holders
WHERE token_id = $1 AND holder_id = $2;