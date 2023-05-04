-- name: PaginatedHolders :many
SELECT * FROM Holders
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: HolderByID :one
SELECT * FROM Holders
WHERE id = $1
LIMIT 1;

-- name: CreateHolder :execresult
INSERT INTO Holders (id)
VALUES ($1);

-- name: DeleteHolder :execresult
DELETE FROM Holders
WHERE id = $1;

-- name: TokenHoldersPaginated :many
SELECT * FROM TokenHolders
ORDER BY token_id, holder_id, block_id
LIMIT $1 OFFSET $2;

-- name: TokensByHolderID :many
SELECT Tokens.*
FROM Tokens
JOIN TokenHolders ON Tokens.id = TokenHolders.token_id
WHERE TokenHolders.holder_id = $1
LIMIT $2 OFFSET $3;

-- name: TokensByHolderIDAndBlockID :many
SELECT Tokens.*
FROM Tokens
JOIN TokenHolders ON Tokens.id = TokenHolders.token_id
WHERE TokenHolders.holder_id = $1 AND TokenHolders.block_id = $2
LIMIT $3 OFFSET $4;

-- name: TokenHoldersByTokenID :many
SELECT Holders.*, TokenHolders.balance
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = $1
LIMIT $2 OFFSET $3;

-- name: TokenHoldersByTokenIDAndBlockID :many
SELECT Holders.*
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = $1 AND TokenHolders.block_id = $2
LIMIT $3 OFFSET $4;

-- name: TokenHoldersByTokenIDAndMinBalance :many
SELECT Holders.*
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = $1 AND TokenHolders.balance >= $2
LIMIT $3 OFFSET $4;

-- name: TokenHoldersByTokenIDAndBlockIDAndMinBalance :many
SELECT Holders.*
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = $1 AND TokenHolders.balance >= $2 AND TokenHolders.block_id = $3
LIMIT $4 OFFSET $5;

-- name: TokenHolderByTokenIDAndHolderID :one
SELECT Holders.*, TokenHolders.balance
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = $1 AND TokenHolders.holder_id = $2;

-- name: TokenHolderByTokenIDAndBlockIDAndHolderID :one
SELECT Holders.*, TokenHolders.balance
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = $1 AND TokenHolders.holder_id = $2 AND TokenHolders.block_id = $3;

-- name: LastBlockByTokenID :one
SELECT block_id 
FROM TokenHolders
WHERE token_id = $1
ORDER BY block_id DESC
LIMIT 1;

-- name: CountTokenHoldersByTokenID :one
SELECT COUNT(holder_id) 
FROM TokenHolders
WHERE token_id = $1;

-- name: CreateTokenHolder :execresult
INSERT INTO TokenHolders (
    token_id,
    holder_id,
    balance,
    block_id
)
VALUES (
    $1, $2, $3, $4
);

-- name: UpdateTokenHolder :execresult
UPDATE TokenHolders
SET balance = sqlc.arg(balance),
    block_id = sqlc.arg(block_id)
WHERE token_id = sqlc.arg(token_id) AND holder_id = sqlc.arg(holder_id) AND block_id = sqlc.arg(block_id);

-- name: DeleteTokenHolder :execresult
DELETE FROM TokenHolders
WHERE token_id = $1 AND holder_id = $2;