-- name: ListCensuses :many
SELECT * FROM censuses
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: CensusByID :one
SELECT * FROM censuses
WHERE id = $1
LIMIT 1;

-- name: CensusByStrategyID :many
SELECT * FROM censuses
WHERE strategy_id = $1;

-- name: CensusByMerkleRoot :one
SELECT * FROM censuses
WHERE merkle_root = $1
LIMIT 1;

-- name: CensusesByStrategyIDAndBlockID :many
SELECT c.* FROM censuses c
JOIN census_blocks cb ON c.id = cb.census_id
WHERE c.strategy_id = sqlc.arg(strategy_id) AND cb.block_id = sqlc.arg(block_id);

-- name: CensusByURI :one
SELECT * FROM censuses
WHERE uri = $1
LIMIT 1;

-- name: CensusesByTokenID :many
SELECT c.* FROM censuses AS c
JOIN strategy_tokens AS st ON c.strategy_id = st.strategy_id
WHERE st.token_id = sqlc.arg(token_id);

-- Get all census paginated for a given token type id
SELECT c.* FROM censuses AS c
JOIN strategy_tokens AS st ON c.strategy_id = st.strategy_id
JOIN tokens AS t ON st.token_id = t.id
JOIN token_types AS tt ON t.type_id = tt.id
WHERE tt.type_name = $1;

-- name: LastCensusID :one
SELECT strategy_id 
FROM censuses 
ORDER BY strategy_id DESC
LIMIT 1;

-- name: CreateCensus :execresult
INSERT INTO censuses (
    id,
    strategy_id,
    merkle_root,
    uri
)
VALUES (
    $1, $2, $3, $4
);

-- name: DeleteCensus :execresult
DELETE FROM censuses
WHERE id = $1;

-- name: UpdateCensus :execresult
UPDATE censuses
SET strategy_id = sqlc.arg(strategy_id),
    merkle_root = sqlc.arg(merkle_root),
    uri = sqlc.arg(uri)
WHERE id = sqlc.arg(id);

-- name: CreateCensusBlock :execresult
INSERT INTO census_blocks (
    census_id,
    block_id
)
VALUES (
    $1, $2
);

-- name: DeleteCensusBlock :execresult
DELETE FROM census_blocks
WHERE census_id = $1 AND block_id = $2;

-- name: UpdateCensusBlock :execresult
UPDATE census_blocks
SET census_id = sqlc.arg(census_id),
    block_id = sqlc.arg(block_id)
WHERE census_id = sqlc.arg(census_id) AND block_id = sqlc.arg(block_id);