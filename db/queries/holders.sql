-- name: ListHolders :many
SELECT * FROM holders
ORDER BY id;

-- name: HolderByID :one
SELECT * FROM holders
WHERE id = ?
LIMIT 1;

-- name: CreateHolder :execresult
INSERT INTO holders (id)
VALUES (?);

-- name: DeleteHolder :execresult
DELETE FROM holders
WHERE id = ?;

-- name: ListTokenHolders :many
SELECT * FROM token_holders
ORDER BY token_id, holder_id, block_id;

-- name: TokensByHolderID :many
SELECT tokens.*
FROM Tokens
JOIN token_holders ON tokens.id = token_holders.token_id
WHERE token_holders.holder_id = ?;

-- name: TokensByHolderIDAndBlockID :many
SELECT tokens.*
FROM Tokens
JOIN token_holders ON tokens.id = token_holders.token_id
WHERE token_holders.holder_id = ? AND token_holders.block_id = ?;

-- name: TokenHoldersByTokenID :many
SELECT holders.*, token_holders.balance
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = ?;

-- name: TokenHoldersByTokenIDAndBlockID :many
SELECT holders.*
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = ? AND token_holders.block_id = ?;

-- name: TokenHoldersByTokenIDAndMinBalance :many
SELECT holders.*
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = ? AND token_holders.balance >= ?;

-- name: TokenHoldersByTokenIDAndBlockIDAndMinBalance :many
SELECT holders.*
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = ? AND token_holders.balance >= ? AND token_holders.block_id = ?;

-- name: TokenHoldersByTokenIDAndExternalID :many
SELECT holders.*, token_holders.balance
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = ? AND token_holders.external_id = ?;

-- name: TokenHolderByTokenIDAndHolderID :one
SELECT holders.*, token_holders.*
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = ? AND token_holders.holder_id = ?;

-- name: TokenHolderByTokenIDAndHolderIDandExternalID :one
SELECT holders.*, token_holders.*
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = ? AND token_holders.holder_id = ? AND token_holders.external_id = ?;

-- name: TokenHolderByTokenIDAndBlockIDAndHolderID :one
SELECT holders.*, token_holders.balance
FROM holders
JOIN token_holders ON holders.id = token_holders.holder_id
WHERE token_holders.token_id = ? AND token_holders.holder_id = ? AND token_holders.block_id = ?;

-- name: LastBlockByTokenID :one
SELECT block_id 
FROM token_holders
WHERE token_id = ?
ORDER BY block_id DESC
LIMIT 1;

-- name: CountTokenHoldersByTokenID :one
SELECT COUNT(holder_id) 
FROM token_holders
WHERE token_id = ?;

-- name: CreateTokenHolder :execresult
INSERT INTO token_holders (
    token_id,
    holder_id,
    balance,
    block_id,
    external_id
)
VALUES (
    ?, ?, ?, ?, ?
);

-- name: UpdateTokenHolderBalance :execresult
UPDATE token_holders
SET balance = sqlc.arg(balance),
    block_id = sqlc.arg(new_block_id)
WHERE token_id = sqlc.arg(token_id) AND holder_id = sqlc.arg(holder_id) AND block_id = sqlc.arg(block_id) AND external_id = sqlc.arg(external_id);

-- name: DeleteTokenHolder :execresult
DELETE FROM token_holders
WHERE token_id = ? AND holder_id = ? AND external_id = ?;