-- name: PaginatedBlocks :many
SELECT * FROM Blocks
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: BlockByID :one
SELECT * FROM Blocks
WHERE id = $1
LIMIT 1;

-- name: BlockByTimestamp :one
SELECT * FROM Blocks
WHERE timestamp = $1
LIMIT 1;

-- name: BlockByRootHash :one
SELECT * FROM Blocks
WHERE root_hash = $1
LIMIT 1;

-- name: CreateBlock :execresult
INSERT INTO Blocks (
    id,
    timestamp,
    root_hash
)
VALUES (
    $1, $2, $3
);

-- name: DeleteBlock :execresult
DELETE FROM Blocks
WHERE id = $1;

-- name: UpdateBlock :execresult
UPDATE Blocks
SET timestamp = sqlc.arg(timestamp),
    root_hash = sqlc.arg(root_hash)
WHERE id = sqlc.arg(id);

-- name: LastBlock :one
SELECT id FROM Blocks 
ORDER BY id DESC 
LIMIT 1;
