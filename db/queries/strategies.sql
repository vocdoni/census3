-- name: ListStrategies :many
SELECT * FROM strategies
ORDER BY id;

-- name: StrategyByID :one
SELECT * FROM strategies
WHERE id = ?
LIMIT 1;

-- name: StrategiesByTokenIDAndChainIDAndExternalID :many
SELECT s.* FROM strategies s
JOIN strategy_tokens st ON st.strategy_id = s.id
WHERE st.token_id = ? AND st.chain_id = ? AND st.external_id = ?
ORDER BY s.id;

-- name: CreateStategy :execresult
INSERT INTO strategies (alias, predicate, uri)
VALUES (?, ?, ?);

-- name: UpdateStrategyIPFSUri :execresult
UPDATE strategies SET uri = ? WHERE id = ?;

-- name: CreateStrategyToken :execresult
INSERT INTO strategy_tokens (
    strategy_id,
    token_id,
    chain_id,
    min_balance,
    external_id
)
VALUES (
    ?, ?, ?, ?, ?
);

-- name: StrategyTokens :many
SELECT *
FROM strategy_tokens
ORDER BY strategy_id, token_id;

-- name: StrategyTokensByStrategyID :many
SELECT st.token_id as id, st.min_balance, t.symbol, t.chain_address, t.chain_id, t.external_id
FROM strategy_tokens st
JOIN tokens t ON t.id = st.token_id
WHERE strategy_id = ?
ORDER BY strategy_id, token_id;