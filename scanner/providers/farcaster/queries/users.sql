-- name: CreateUser :execresult
INSERT INTO users (
    fid,
    username,
    signer,
    custody_address,
    app_keys,
    recovery_address)
VALUES (?, ?, ?, ?, ?, ?);

-- name: DeleteUser :execresult
DELETE FROM users WHERE fid = ?;

-- name: GetUserByFID :one
SELECT * FROM users WHERE fid = ?;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = ?;

-- name: GetUserBySigner :one
SELECT * FROM users WHERE signer = ?;

-- name: ListUsers :many
SELECT * FROM users ORDER BY fid ASC;

-- name: CountUsers :one
SELECT COUNT(*) FROM users;

-- name: CountUserAppKeys :one
SELECT COUNT(app_keys) FROM users WHERE fid = ?;

-- name: UpdateUser :execresult
UPDATE users 
SET username = sqlc.arg(username),
    signer = sqlc.arg(signer),
    custody_address = sqlc.arg(custody_address),
    app_keys = sqlc.arg(app_keys),
    recovery_address = sqlc.arg(recovery_address)
WHERE fid = ?;

-- name: UpdateCustodyAddress :execresult
UPDATE users
SET custody_address = sqlc.arg(custody_address)
WHERE fid = ?;

-- name: UpdateUserAppKeys :execresult
UPDATE users
SET app_keys = sqlc.arg(app_keys)
WHERE fid = ?;

-- name: UpdateUserSigner :execresult
UPDATE users
SET signer = sqlc.arg(signer)
WHERE fid = ?;

-- name: UpdateUserRecoveryAddress :execresult
UPDATE users
SET recovery_address = sqlc.arg(recovery_address)
WHERE fid = ?;

-- name: UpdateUsername :execresult
UPDATE users
SET username = sqlc.arg(username)
WHERE fid = ?;