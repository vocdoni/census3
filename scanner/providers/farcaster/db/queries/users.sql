-- name: CreateUser :execresult
INSERT INTO users (
    fid,
    username,
    signer,
    custodyAddress,
    appkeys,
    recoveryAddress)
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
SELECT COUNT(appkeys) FROM users WHERE fid = ?;

-- name: UpdateUser :execresult
UPDATE users 
SET username = sqlc.arg(username),
    signer = sqlc.arg(signer),
    custodyAddress = sqlc.arg(custodyAddress),
    appkeys = sqlc.arg(appkeys),
    recoveryAddress = sqlc.arg(recoveryAddress)
WHERE fid = ?;

-- name: UpdateCustodyAddress :execresult
UPDATE users
SET custodyAddress = sqlc.arg(custodyAddress)
WHERE fid = ?;

-- name: UpdateUserAppKeys :execresult
UPDATE users
SET appkeys = sqlc.arg(appkeys)
WHERE fid = ?;

-- name: UpdateUserRecoveryAddress :execresult
UPDATE users
SET recoveryAddress = sqlc.arg(recoveryAddress)
WHERE fid = ?;

-- name: UpdateUsername :execresult
UPDATE users
SET username = sqlc.arg(username)
WHERE fid = ?;