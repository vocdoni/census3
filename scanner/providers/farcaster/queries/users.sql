-- name: CreateUser :execresult
INSERT INTO users (fid) VALUES (?);

-- name: DeleteUser :execresult
DELETE FROM users WHERE fid = ?;

-- name: GetUserByFID :one
SELECT * FROM users WHERE fid = ?;

-- name: ListUsers :many
SELECT * FROM users ORDER BY fid ASC;

-- name: CountUsers :one
SELECT COUNT(*) FROM users;
