-- name: ListCensuses :many
SELECT * FROM censuses
ORDER BY id;

-- name: CensusByID :one
SELECT * FROM censuses
WHERE id = ?
LIMIT 1;

-- name: CensusByStrategyID :many
SELECT * FROM censuses
WHERE strategy_id = ?;

-- name: CensusByMerkleRoot :one
SELECT * FROM censuses
WHERE merkle_root = ?
LIMIT 1;

-- name: CensusByQueueID :one
SELECT * FROM censuses
WHERE queue_id = ?
LIMIT 1;

-- name: CensusByURI :one
SELECT * FROM censuses
WHERE uri = ?
LIMIT 1;

-- name: CensusesByTokenID :many
SELECT c.* FROM censuses AS c
JOIN strategy_tokens AS st ON c.strategy_id = st.strategy_id
WHERE st.token_id = sqlc.arg(token_id)
LIMIT ? OFFSET ?;

-- name: CensusesByTokenType :many
SELECT c.* FROM censuses AS c
JOIN strategy_tokens AS st ON c.strategy_id = st.strategy_id
JOIN tokens AS t ON st.token_id = t.id
JOIN token_types AS tt ON t.type_id = tt.id
WHERE tt.type_name = ?;

-- name: LastCensusID :one
SELECT id 
FROM censuses 
ORDER BY id DESC
LIMIT 1;

-- name: CreateCensus :execresult
INSERT INTO censuses (
    id,
    strategy_id,
    merkle_root,
    uri,
    size, 
    weight,
    census_type,
    queue_id
)
VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: DeleteCensus :execresult
DELETE FROM censuses
WHERE id = ?;

-- name: UpdateCensus :execresult
UPDATE censuses
SET merkle_root = sqlc.arg(merkle_root),
    uri = sqlc.arg(uri),
    size = sqlc.arg(size),
    weight = sqlc.arg(weight)
WHERE id = sqlc.arg(id);
