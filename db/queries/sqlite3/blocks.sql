-- name: PaginatedBlocks :many
SELECT * FROM Blocks
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: BlockByID :one
SELECT * FROM Blocks
WHERE id = ?
LIMIT 1;

-- name: BlockByTimestamp :one
SELECT * FROM Blocks
WHERE timestamp = ?
LIMIT 1;

-- name: BlockByRootHash :one
SELECT * FROM Blocks
WHERE root_hash = ?
LIMIT 1;

-- name: CreateBlock :execresult
INSERT INTO Blocks (
    id,
    timestamp,
    root_hash
)
VALUES (
    ?, ?, ?
);

-- name: DeleteBlock :execresult
DELETE FROM Blocks
WHERE id = ?;

-- name: UpdateBlock :execresult
UPDATE Blocks
SET timestamp = sqlc.arg(timestamp),
    root_hash = sqlc.arg(root_hash)
WHERE id = sqlc.arg(id);

-- name: LastBlock :one
SELECT id FROM Blocks 
ORDER BY id DESC 
LIMIT 1;
