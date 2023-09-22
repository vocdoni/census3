-- name: BlockByID :one
SELECT * FROM blocks
WHERE id = ?
LIMIT 1;

-- name: CreateBlock :execresult
INSERT INTO blocks (
    id,
    timestamp,
    root_hash
)
VALUES (
    ?, ?, ?
);

-- name: LastBlock :one
SELECT id FROM blocks 
ORDER BY id DESC 
LIMIT 1;