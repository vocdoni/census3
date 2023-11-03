-- name: CensusByID :one
SELECT * FROM censuses
WHERE id = ?
LIMIT 1;

-- name: CensusByStrategyID :many
SELECT * FROM censuses
WHERE strategy_id = ?;


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