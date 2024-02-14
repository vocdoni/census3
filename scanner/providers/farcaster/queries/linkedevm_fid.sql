-- name: CreateLinkedEVMFID :execresult
INSERT INTO linkedevm_fid (fid, evm_address) VALUES (?, ?);

-- name: DeleteLinkedEVMFID :execresult
DELETE FROM linkedevm_fid WHERE fid = ? AND evm_address = ?;

-- name: GetLinkedEVMFID :many
SELECT evm_address FROM linkedevm_fid WHERE fid = ?;

-- name: CheckLinkedEVMForAny :one
SELECT EXISTS(SELECT 1 FROM linkedevm_fid WHERE evm_address = ?);

-- name: ListUsersWithLinkedEVM :many
SELECT u.fid, u.signer, l.evm_address FROM users u
JOIN linkedevm_fid l ON u.fid = l.fid;

-- name: ListUsersNotLinkedEVM :many
SELECT * FROM users WHERE fid NOT IN (SELECT fid FROM linkedevm_fid);

-- name: GetFIDByLinkedEVM :many
SELECT fid FROM linkedevm_fid WHERE evm_address = ?;

-- name: GetFidsByLinkedEVM :many
SELECT users.fid, users.signer
FROM users
JOIN linkedevm_fid ON users.fid = linkedevm_fid.fid
WHERE linkedevm_fid.evm_address = ?;