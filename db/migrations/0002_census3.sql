-- +goose Up

-- stategies table schema updates
ALTER TABLE strategies ADD COLUMN alias TEXT;

-- tokens table schema updates
CREATE TABLE tokens_copy (
    id BLOB NOT NULL,
    name TEXT,
    symbol TEXT,
    decimals INTEGER,
    total_supply BLOB,
    creation_block BIGINT,
    type_id INTEGER NOT NULL,
    synced BOOLEAN NOT NULL,
    tags TEXT,
    chain_id INTEGER NOT NULL,
    PRIMARY KEY (id, chain_id),
    FOREIGN KEY (type_id) REFERENCES token_types(id) ON DELETE CASCADE
);
INSERT INTO tokens_copy SELECT * FROM tokens;
DROP TABLE tokens;
DROP INDEX IF EXISTS idx_tokens_type_id;
ALTER TABLE tokens_copy RENAME TO tokens;
CREATE INDEX idx_tokens_type_id ON tokens(type_id);

-- token_holders table schema updates
ALTER TABLE token_holders ADD COLUMN chain_id INTEGER;
CREATE TABLE token_holders_copy (
    token_id BLOB NOT NULL,
    holder_id BLOB NOT NULL,
    balance BLOB NOT NULL,
    block_id INTEGER NOT NULL,
    chain_id INTEGER NOT NULL,
    PRIMARY KEY (token_id, holder_id, block_id, chain_id),
    FOREIGN KEY (token_id) REFERENCES tokens(id) ON DELETE CASCADE,
    FOREIGN KEY (holder_id) REFERENCES holders(id) ON DELETE CASCADE,
    FOREIGN KEY (block_id) REFERENCES blocks(id) ON DELETE CASCADE
);
INSERT INTO token_holders_copy (token_id, holder_id, balance, block_id, chain_id) 
SELECT * FROM (
    SELECT token_id, holder_id, balance, block_id, (
        SELECT token.chain_id FROM tokens AS token WHERE token.id = token_holders.token_id
    ) AS chain_id FROM token_holders
);
DROP INDEX IF EXISTS idx_token_holders_token_id;
DROP INDEX IF EXISTS idx_token_holders_holder_id;
DROP INDEX IF EXISTS idx_token_holders_block_id;
DROP TABLE token_holders;
ALTER TABLE token_holders_copy RENAME TO token_holders;
CREATE INDEX idx_token_holders_token_id ON token_holders(token_id);
CREATE INDEX idx_token_holders_holder_id ON token_holders(holder_id);
CREATE INDEX idx_token_holders_block_id ON token_holders(block_id);

-- strategy tokens schema updates
CREATE TABLE strategy_tokens_copy (
    strategy_id INTEGER NOT NULL,
    token_id BLOB NOT NULL,
    min_balance BLOB NOT NULL,
    chain_id INTEGER NOT NULL,
    PRIMARY KEY (strategy_id, token_id),
    FOREIGN KEY (strategy_id) REFERENCES strategies(id) ON DELETE CASCADE,
    FOREIGN KEY (token_id) REFERENCES tokens(id) ON DELETE CASCADE
);
INSERT INTO strategy_tokens_copy (strategy_id, token_id, min_balance, chain_id) 
SELECT * FROM (
    SELECT strategy_id, token_id, min_balance, (
        SELECT token.chain_id FROM tokens AS token WHERE token.id = strategy_tokens.token_id
    ) AS chain_id FROM strategy_tokens
);

DROP INDEX IF EXISTS idx_strategy_tokens_strategy_id;
DROP INDEX IF EXISTS idx_strategy_tokens_token_id;
DROP TABLE strategy_tokens;
ALTER TABLE strategy_tokens_copy RENAME TO strategy_tokens;
CREATE INDEX idx_strategy_tokens_strategy_id ON strategy_tokens(strategy_id);
CREATE INDEX idx_strategy_tokens_token_id ON strategy_tokens(token_id);