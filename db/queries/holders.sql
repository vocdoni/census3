-- name: PaginatedHolders :many
SELECT * FROM Holders
ORDER BY id
LIMIT ? OFFSET ?;

-- name: HolderByID :one
SELECT * FROM Holders
WHERE id = ?
LIMIT 1;

-- name: CreateHolder :execresult
INSERT INTO Holders (id)
VALUES (?);

-- name: DeleteHolder :execresult
DELETE FROM Holders
WHERE id = ?;

-- name: TokenHoldersPaginated :many
SELECT * FROM TokenHolders
ORDER BY token_id, holder_id, block_id
LIMIT ? OFFSET ?;

-- name: TokensByHolderID :many
SELECT Tokens.*
FROM Tokens
JOIN TokenHolders ON Tokens.id = TokenHolders.token_id
WHERE TokenHolders.holder_id = ?
LIMIT ? OFFSET ?;

-- name: TokensByHolderIDAndBlockID :many
SELECT Tokens.*
FROM Tokens
JOIN TokenHolders ON Tokens.id = TokenHolders.token_id
WHERE TokenHolders.holder_id = ? AND TokenHolders.block_id = ?
LIMIT ? OFFSET ?;

-- name: TokenHoldersByTokenID :many
SELECT Holders.*
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = ?
LIMIT ? OFFSET ?;

-- name: TokenHoldersByTokenIDAndBlockID :many
SELECT Holders.*
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = ? AND TokenHolders.block_id = ?
LIMIT ? OFFSET ?;

-- name: TokenHoldersByTokenIDAndMinBalance :many
SELECT Holders.*
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = ? AND TokenHolders.balance >= ?
LIMIT ? OFFSET ?;

-- name: TokenHoldersByTokenIDAndBlockIDAndMinBalance :many
SELECT Holders.*
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = ? AND TokenHolders.balance >= ? AND TokenHolders.block_id = ?
LIMIT ? OFFSET ?;

-- name: TokenHolderByTokenIDAndHolderID :one
SELECT Holders.*, TokenHolders.balance
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = ? AND TokenHolders.holder_id = ?;

-- name: TokenHolderByTokenIDAndBlockIDAndHolderID :one
SELECT Holders.*, TokenHolders.balance
FROM Holders
JOIN TokenHolders ON Holders.id = TokenHolders.holder_id
WHERE TokenHolders.token_id = ? AND TokenHolders.holder_id = ? AND TokenHolders.block_id = ?;

-- name: LastBlockByTokenID :one
SELECT block_id 
FROM TokenHolders
WHERE token_id = ?
ORDER BY block_id DESC
LIMIT 1;

-- name: CreateTokenHolder :execresult
INSERT INTO TokenHolders (
    token_id,
    holder_id,
    balance,
    block_id
)
VALUES (
    ?, ?, ?, ?
);

-- name: UpdateTokenHolder :execresult
UPDATE TokenHolders
SET balance = sqlc.arg(balance),
    block_id = sqlc.arg(block_id)
WHERE token_id = sqlc.arg(token_id) AND holder_id = sqlc.arg(holder_id) AND block_id = sqlc.arg(block_id);

-- name: DeleteTokenHolder :execresult
DELETE FROM TokenHolders
WHERE token_id = ? AND holder_id = ? AND block_id = ?;
