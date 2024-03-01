-- name: LastBlock :one
SELECT block_number FROM latest_blocks WHERE contract = sqlc.arg(contract);

-- name: SetLastBlock :execresult
UPDATE latest_blocks SET block_number = sqlc.arg(block_number) WHERE contract = sqlc.arg(contract);

-- name: ExistsLatestBlock :one
SELECT EXISTS (SELECT block_number FROM latest_blocks WHERE contract = sqlc.arg(contract));

-- name: InsertLatestBlock :execresult
INSERT INTO latest_blocks (contract, block_number) VALUES (sqlc.arg(contract), sqlc.arg(block_number));
