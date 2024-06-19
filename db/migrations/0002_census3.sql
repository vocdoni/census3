-- +goose Up
ALTER TABLE strategy_tokens ADD COLUMN token_alias TEXT NOT NULL DEFAULT '';

CREATE INDEX idx_strategy_tokens_token_alias ON strategy_tokens(token_alias);

UPDATE strategy_tokens
SET token_alias = (
    SELECT t.symbol
    FROM tokens t
    WHERE strategy_tokens.token_id = t.id
    AND strategy_tokens.chain_id = t.chain_id
    AND (strategy_tokens.external_id = t.external_id OR strategy_tokens.external_id IS NULL AND t.external_id IS NULL)
)
WHERE token_alias = '';

-- +goose Down
DROP INDEX IF EXISTS idx_strategy_tokens_token_alias;

CREATE TABLE strategy_tokens_temp AS
SELECT strategy_id, token_id, min_balance, chain_id, external_id
FROM strategy_tokens;

DROP TABLE strategy_tokens;

ALTER TABLE strategy_tokens_temp RENAME TO strategy_tokens;