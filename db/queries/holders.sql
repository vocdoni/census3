
-- name: ListTokenHolders :many
SELECT *
FROM token_holders
WHERE token_id = ? AND chain_id = ? AND external_id = ?;

-- name: GetTokenHolder :one
SELECT *
FROM token_holders
WHERE token_id = ? 
    AND holder_id = ? 
    AND chain_id = ?
    AND external_id = ?
    AND balance > '0';

-- name: ExistTokenHolder :one
SELECT EXISTS (
    SELECT holder_id 
    FROM token_holders
    WHERE token_id = ? 
        AND holder_id = ?
        AND chain_id = ?
        AND external_id = ?
);

-- name: CountTokenHolders :one
SELECT COUNT(holder_id) 
FROM token_holders
WHERE token_id = ?
    AND chain_id = ?
    AND external_id = ?
    AND balance >= ?;

-- name: CreateTokenHolder :execresult
INSERT INTO token_holders (
    token_id,
    holder_id,
    balance,
    block_id,
    chain_id,
    external_id
)
VALUES (
    ?, ?, ?, ?, ?, ?
);

-- name: UpdateTokenHolderBalance :execresult
UPDATE token_holders
SET balance = sqlc.arg(balance),
    block_id = sqlc.arg(new_block_id)
WHERE token_id = sqlc.arg(token_id) 
    AND holder_id = sqlc.arg(holder_id) 
    AND block_id = sqlc.arg(block_id) 
    AND chain_id = sqlc.arg(chain_id) 
    AND external_id = sqlc.arg(external_id);

-- name: TokenHoldersByTokenIDAndChainIDAndMinBalance :many
SELECT token_holders.holder_id, token_holders.balance
FROM token_holders
WHERE token_holders.token_id = ? 
    AND token_holders.chain_id = ?
    AND token_holders.external_id = ?
    AND token_holders.balance >= ?;

-- name: DeleteTokenHolder :execresult
DELETE FROM token_holders WHERE token_id = ? AND chain_id = ? AND external_id = ?;

-- name: TokenHoldersByStrategyID :many
SELECT token_holders.holder_id, token_holders.balance, strategy_tokens.min_balance
FROM token_holders
JOIN strategy_tokens 
    ON strategy_tokens.token_id = token_holders.token_id
    AND strategy_tokens.chain_id = token_holders.chain_id
    AND strategy_tokens.external_id = token_holders.external_id
WHERE strategy_tokens.strategy_id = ?
    AND strategy_tokens.min_balance <= token_holders.balance;

-- name: ANDOperator :many
;WITH holders_a as (
    SELECT th.holder_id, th.balance
    FROM token_holders th
    WHERE th.token_id = sqlc.arg(token_id_a) 
        AND th.chain_id = sqlc.arg(chain_id_a)
        AND th.external_id = sqlc.arg(external_id_a)
        AND th.balance >= sqlc.arg(min_balance_a)
),
holders_b as (
    SELECT th.holder_id, th.balance
    FROM token_holders th
    WHERE th.token_id = sqlc.arg(token_id_b) 
        AND th.chain_id = sqlc.arg(chain_id_b)
        AND th.external_id = sqlc.arg(external_id_b)
        AND th.balance >= sqlc.arg(min_balance_b)
)
SELECT holders_a.holder_id, IFNULL(holders_a.balance, '0') as balance_a, IFNULL(holders_b.balance, '0') as balance_b
FROM holders_a
INNER JOIN holders_b ON holders_a.holder_id = holders_b.holder_id;

-- name: OROperator :many
SELECT holder_ids.holder_id, IFNULL(a.balance, '0') AS balance_a, IFNULL(b.balance, '0') AS balance_b
FROM (
    SELECT th.holder_id
    FROM token_holders th
    WHERE (
        th.token_id = sqlc.arg(token_id_a) 
        AND th.chain_id = sqlc.arg(chain_id_a)
        AND th.external_id = sqlc.arg(external_id_a)
        AND th.balance >= sqlc.arg(min_balance_a)
    ) OR (
        th.token_id = sqlc.arg(token_id_b) 
        AND th.chain_id = sqlc.arg(chain_id_b)
        AND th.external_id = sqlc.arg(external_id_b)
        AND th.balance >= sqlc.arg(min_balance_b)
    )
) as holder_ids
LEFT JOIN (
    SELECT th_b.holder_id, th_b.balance
    FROM token_holders th_b
    WHERE th_b.token_id = sqlc.arg(token_id_a) 
        AND th_b.chain_id = sqlc.arg(chain_id_a)
        AND th_b.external_id = sqlc.arg(external_id_a)
        AND th_b.balance >= sqlc.arg(min_balance_a)
) AS a ON holder_ids.holder_id = a.holder_id
LEFT JOIN (
    SELECT th_a.holder_id, th_a.balance
    FROM token_holders th_a
    WHERE th_a.token_id = sqlc.arg(token_id_b) 
        AND th_a.chain_id = sqlc.arg(chain_id_b)
        AND th_a.external_id = sqlc.arg(external_id_b)
        AND th_a.balance >= sqlc.arg(min_balance_b)
) AS b ON holder_ids.holder_id = b.holder_id
GROUP BY holder_ids.holder_id;
