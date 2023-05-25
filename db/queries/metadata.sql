-- name: ChainID :one
SELECT chainID 
FROM metadata
LIMIT 1;

-- name: SetChainID :execresult
INSERT INTO metadata (
    chainID
)
VALUES (
    ?
);