-- +goose Up

-- strategy tokens schema updates
ALTER TABLE strategy_tokens ADD COLUMN external_id BLOB NOT NULL DEFAULT '';
CREATE TABLE strategy_tokens_copy (
    strategy_id INTEGER NOT NULL,
    token_id BLOB NOT NULL,
    min_balance BLOB NOT NULL,
    chain_id INTEGER NOT NULL,
    external_id TEXT NOT NULL DEFAULT '',
    PRIMARY KEY (strategy_id, token_id, chain_id, external_id),
    FOREIGN KEY (strategy_id) REFERENCES strategies(id) ON DELETE CASCADE,
    FOREIGN KEY (token_id) REFERENCES tokens(id) ON DELETE CASCADE
);
INSERT INTO strategy_tokens_copy (strategy_id, token_id, min_balance, chain_id, external_id) 
SELECT * FROM (
    SELECT strategy_id, token_id, min_balance, external_id, (
        SELECT token.chain_id FROM tokens AS token WHERE token.id = strategy_tokens.token_id
    ) AS chain_id FROM strategy_tokens
);
DROP TABLE strategy_tokens;
ALTER TABLE strategy_tokens_copy RENAME TO strategy_tokens;
CREATE INDEX idx_strategy_tokens_strategy_id ON strategy_tokens(strategy_id);
CREATE INDEX idx_strategy_tokens_token_id ON strategy_tokens(token_id);

-- tokens table schema updates
ALTER TABLE tokens ADD COLUMN chain_address TEXT NOT NULL DEFAULT '';
ALTER TABLE tokens ADD COLUMN external_id TEXT NOT NULL DEFAULT '';
ALTER TABLE tokens ADD COLUMN icon_uri TEXT NOT NULL DEFAULT '';
ALTER TABLE tokens RENAME COLUMN tag TO tags;
UPDATE tokens SET name = '' WHERE name IS NULL;
UPDATE tokens SET symbol = '' WHERE symbol IS NULL;
UPDATE tokens SET decimals = 0 WHERE decimals IS NULL;
UPDATE tokens SET total_supply = '' WHERE total_supply IS NULL;
UPDATE tokens SET creation_block = 0 WHERE creation_block IS NULL;
UPDATE tokens SET tags = '' WHERE tags IS NULL;
UPDATE tokens SET chain_address = '' WHERE chain_address IS NULL;
UPDATE tokens SET external_id = '' WHERE external_id IS NULL;
CREATE TABLE tokens_copy (
    id BLOB NOT NULL,
    name TEXT NOT NULL DEFAULT '',
    symbol TEXT NOT NULL DEFAULT '',
    decimals INTEGER NOT NULL DEFAULT 0,
    total_supply BLOB NOT NULL DEFAULT '',
    creation_block BIGINT NOT NULL DEFAULT 0,
    type_id INTEGER NOT NULL DEFAULT 0,
    synced BOOLEAN NOT NULL DEFAULT 0,
    tags TEXT NOT NULL DEFAULT '',
    chain_id INTEGER NOT NULL DEFAULT 0,
    chain_address TEXT NOT NULL DEFAULT '',
    external_id TEXT NULL DEFAULT '',
    default_strategy INTEGER NOT NULL DEFAULT 0,
    icon_uri TEXT NOT NULL DEFAULT '',
    PRIMARY KEY (id, chain_id, external_id),
    FOREIGN KEY (type_id) REFERENCES token_types(id) ON DELETE CASCADE,
    FOREIGN KEY (default_strategy) REFERENCES strategies(id) ON DELETE CASCADE
);
INSERT INTO tokens_copy (id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id, chain_address, external_id, default_strategy, icon_uri) 
SELECT * FROM (
    SELECT id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id, chain_address, external_id, icon_uri, (
        SELECT strategy_id FROM strategy_tokens WHERE token_id = tokens.id AND strategy_tokens.chain_id = tokens.chain_id AND strategy_tokens.external_id = tokens.external_id LIMIT 1
    ) AS default_strategy FROM tokens
);
DROP TABLE tokens;
ALTER TABLE tokens_copy RENAME TO tokens;
CREATE INDEX idx_tokens_type_id ON tokens(type_id);

-- token_holders table schema updates
ALTER TABLE token_holders ADD COLUMN chain_id INTEGER NOT NULL DEFAULT 0;
ALTER TABLE token_holders ADD COLUMN external_id TEXT NOT NULL DEFAULT '';
CREATE TABLE token_holders_copy (
    token_id BLOB NOT NULL,
    holder_id BLOB NOT NULL,
    balance BLOB NOT NULL,
    block_id INTEGER NOT NULL,
    chain_id INTEGER NOT NULL,
    external_id TEXT NULL DEFAULT '',
    PRIMARY KEY (token_id, holder_id, block_id, chain_id, external_id),
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
DROP TABLE token_holders;
ALTER TABLE token_holders_copy RENAME TO token_holders;
CREATE INDEX idx_token_holders_token_id ON token_holders(token_id);
CREATE INDEX idx_token_holders_holder_id ON token_holders(holder_id);
CREATE INDEX idx_token_holders_block_id ON token_holders(block_id);

-- stategies table schema updates
ALTER TABLE strategies ADD COLUMN alias TEXT NOT NULL DEFAULT '';
ALTER TABLE strategies ADD COLUMN uri TEXT NOT NULL DEFAULT '';