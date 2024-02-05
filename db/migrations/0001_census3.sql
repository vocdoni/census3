-- +goose Up
CREATE TABLE strategies (
    id INTEGER PRIMARY KEY,
    predicate TEXT NOT NULL,
    alias TEXT NOT NULL DEFAULT '',
    uri TEXT NOT NULL DEFAULT ''
);

CREATE TABLE token_types (
    id INTEGER PRIMARY KEY,
    type_name TEXT NOT NULL
);

INSERT INTO token_types (id, type_name) VALUES (0, 'unknown');
INSERT INTO token_types (id, type_name) VALUES (1, 'erc20');
INSERT INTO token_types (id, type_name) VALUES (2, 'erc721');;
INSERT INTO token_types (id, type_name) VALUES (3, 'erc777');
INSERT INTO token_types (id, type_name) VALUES (4, 'poap');
INSERT INTO token_types (id, type_name) VALUES (5, 'gitcoinpassport');

CREATE TABLE tokens (
    id BLOB NOT NULL,
    name TEXT NOT NULL DEFAULT '',
    symbol TEXT NOT NULL DEFAULT '',
    decimals INTEGER NOT NULL DEFAULT 0,
    total_supply TEXT NOT NULL DEFAULT '',
    creation_block BIGINT NOT NULL DEFAULT 0,
    type_id INTEGER NOT NULL DEFAULT 0,
    synced BOOLEAN NOT NULL DEFAULT 0,
    tags TEXT NOT NULL DEFAULT '',
    chain_id INTEGER NOT NULL DEFAULT 0,
    chain_address TEXT NOT NULL DEFAULT '',
    external_id TEXT NULL DEFAULT '',
    default_strategy INTEGER NOT NULL DEFAULT 0,
    icon_uri TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_block BIGINT NOT NULL DEFAULT 0,
    analysed_transfers BIGINT NOT NULL DEFAULT 0,
    PRIMARY KEY (id, chain_id, external_id),
    FOREIGN KEY (type_id) REFERENCES token_types(id) ON DELETE CASCADE,
    FOREIGN KEY (default_strategy) REFERENCES strategies(id) ON DELETE CASCADE
);
CREATE INDEX idx_tokens_type_id ON tokens(type_id);

CREATE TABLE censuses (
    id INTEGER PRIMARY KEY,
    strategy_id INTEGER NOT NULL,
    merkle_root BLOB NOT NULL,
    uri TEXT,
    size INTEGER,
    weight BLOB,
    census_type INTEGER NOT NULL,
    queue_id TEXT NOT NULL,
    accuracy FLOAT NOT NULL DEFAULT '100.0',
    FOREIGN KEY (strategy_id) REFERENCES strategies(id) ON DELETE CASCADE,
    UNIQUE(id, merkle_root)
);
CREATE INDEX idx_censuses_strategy_id ON censuses(strategy_id);

CREATE TABLE token_holders (
    token_id BLOB NOT NULL,
    holder_id BLOB NOT NULL,
    balance TEXT NOT NULL,
    block_id INTEGER NOT NULL,
    chain_id INTEGER NOT NULL,
    external_id TEXT NULL DEFAULT '',
    PRIMARY KEY (token_id, holder_id, block_id, chain_id, external_id),
    FOREIGN KEY (token_id) REFERENCES tokens(id) ON DELETE CASCADE
);
CREATE INDEX idx_token_holders_token_id ON token_holders(token_id);
CREATE INDEX idx_token_holders_holder_id ON token_holders(holder_id);
CREATE INDEX idx_token_holders_block_id ON token_holders(block_id);

CREATE TABLE strategy_tokens (
    strategy_id INTEGER NOT NULL,
    token_id BLOB NOT NULL,
    min_balance TEXT NOT NULL,
    chain_id INTEGER NOT NULL,
    external_id TEXT NOT NULL DEFAULT '',
    PRIMARY KEY (strategy_id, token_id, chain_id, external_id),
    FOREIGN KEY (strategy_id) REFERENCES strategies(id) ON DELETE CASCADE,
    FOREIGN KEY (token_id) REFERENCES tokens(id) ON DELETE CASCADE
);
CREATE INDEX idx_strategy_tokens_strategy_id ON strategy_tokens(strategy_id);
CREATE INDEX idx_strategy_tokens_token_id ON strategy_tokens(token_id);

-- +goose Down
DROP INDEX IF EXISTS idx_strategy_tokens_token_id;
DROP INDEX IF EXISTS idx_strategy_tokens_strategy_id;
DROP INDEX IF EXISTS idx_token_holders_block_id;
DROP INDEX IF EXISTS idx_token_holders_holder_id;
DROP INDEX IF EXISTS idx_token_holders_token_id;
DROP INDEX IF EXISTS idx_censuses_strategy_id;
DROP INDEX IF EXISTS idx_tokens_type_id;

DROP TABLE IF EXISTS strategy_tokens;
DROP TABLE IF EXISTS token_holders;
DROP TABLE IF EXISTS censuses;
DROP TABLE IF EXISTS tokens;
DROP TABLE IF EXISTS token_types;
DROP TABLE IF EXISTS strategies;