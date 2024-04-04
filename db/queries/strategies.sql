-- name: ListStrategies :many
SELECT * FROM strategies
ORDER BY id;

-- name: NextStrategiesPage :many
SELECT * FROM strategies
WHERE id >= sqlc.arg(page_cursor)
ORDER BY id ASC 
LIMIT ?;

-- name: PrevStrategiesPage :many
SELECT * FROM (
    SELECT * FROM strategies
    WHERE id <= sqlc.arg(page_cursor)
    ORDER BY id DESC 
    LIMIT ?
) as strategy ORDER BY strategy.id ASC; 

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

-- name: ExistsStrategyByURI :one
SELECT EXISTS(SELECT 1 FROM strategies WHERE uri = ?);

-- name: StrategyTokensByStrategyID :many
SELECT st.token_id as id, st.min_balance, t.symbol, t.chain_address, t.chain_id, t.external_id
FROM strategy_tokens st
JOIN tokens t ON t.id = st.token_id AND t.chain_id = st.chain_id AND t.external_id = st.external_id
WHERE st.strategy_id = ?
ORDER BY strategy_id, token_id;

-- name: StrategyTokens :many
SELECT st.token_id, st.min_balance, st.chain_id, st.external_id, t.chain_address, t.symbol, t.icon_uri
FROM strategy_tokens st
JOIN tokens t ON st.token_id = t.id AND st.chain_id = t.chain_id AND st.external_id = t.external_id
WHERE st.strategy_id = ?;

-- name: StrategyTokensContainsType :one
SELECT EXISTS(
    SELECT 1 FROM strategy_tokens st
    JOIN tokens t ON st.token_id = t.id AND st.chain_id = t.chain_id AND st.external_id = t.external_id
    WHERE st.strategy_id = sqlc.arg(strategy_id) AND t.type_id = sqlc.arg(type_id)
);

-- name: DeleteStrategiesByToken :execresult
DELETE FROM strategies WHERE id IN (
    SELECT strategy_id FROM strategy_tokens WHERE token_id = ? AND chain_id = ? AND external_id = ?
);

-- name: DeleteStrategyTokensByToken :execresult
DELETE FROM strategy_tokens WHERE token_id = ? AND chain_id = ? AND external_id = ?;