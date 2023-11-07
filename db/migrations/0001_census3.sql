-- +goose Up
CREATE TABLE strategies (
    id INTEGER PRIMARY KEY,
    predicate TEXT NOT NULL
);

CREATE TABLE token_types (
    id INTEGER PRIMARY KEY,
    type_name TEXT NOT NULL UNIQUE
);

INSERT INTO token_types (type_name) VALUES ('erc20');
INSERT INTO token_types (type_name) VALUES ('erc721');
INSERT INTO token_types (type_name) VALUES ('erc1155');
INSERT INTO token_types (type_name) VALUES ('erc777');
INSERT INTO token_types (type_name) VALUES ('nation3');
INSERT INTO token_types (type_name) VALUES ('want');

CREATE TABLE tokens (
    id BLOB PRIMARY KEY NOT NULL,
    name TEXT,
    symbol TEXT,
    decimals INTEGER,
    total_supply BLOB,
    creation_block BIGINT,
    type_id INTEGER NOT NULL,
    synced BOOLEAN NOT NULL,
    tag TEXT,
    chain_id INTEGER NOT NULL,
    UNIQUE (id, chain_id),
    FOREIGN KEY (type_id) REFERENCES token_types(id) ON DELETE CASCADE
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
    FOREIGN KEY (strategy_id) REFERENCES strategies(id) ON DELETE CASCADE,
    UNIQUE(id, merkle_root)
);
CREATE INDEX idx_censuses_strategy_id ON censuses(strategy_id);

CREATE TABLE blocks (
    id INTEGER PRIMARY KEY NOT NULL,
    timestamp TEXT NOT NULL UNIQUE,
    root_hash BLOB NOT NULL UNIQUE
);

CREATE TABLE holders (
    id BLOB PRIMARY KEY NOT NULL
);

CREATE TABLE token_holders (
    token_id BLOB NOT NULL,
    holder_id BLOB NOT NULL,
    balance BLOB NOT NULL,
    block_id INTEGER NOT NULL,
    PRIMARY KEY (token_id, holder_id, block_id),
    FOREIGN KEY (token_id) REFERENCES tokens(id) ON DELETE CASCADE,
    FOREIGN KEY (holder_id) REFERENCES holders(id) ON DELETE CASCADE,
    FOREIGN KEY (block_id) REFERENCES blocks(id) ON DELETE CASCADE
);
CREATE INDEX idx_token_holders_token_id ON token_holders(token_id);
CREATE INDEX idx_token_holders_holder_id ON token_holders(holder_id);
CREATE INDEX idx_token_holders_block_id ON token_holders(block_id);

CREATE TABLE strategy_tokens (
    strategy_id INTEGER NOT NULL,
    token_id BLOB NOT NULL,
    min_balance BLOB NOT NULL,
    method_hash BLOB NOT NULL,
    PRIMARY KEY (strategy_id, token_id),
    FOREIGN KEY (strategy_id) REFERENCES strategies(id) ON DELETE CASCADE,
    FOREIGN KEY (token_id) REFERENCES tokens(id) ON DELETE CASCADE
);
CREATE INDEX idx_strategy_tokens_strategy_id ON strategy_tokens(strategy_id);
CREATE INDEX idx_strategy_tokens_token_id ON strategy_tokens(token_id);

CREATE TABLE census_blocks (
    census_id INTEGER NOT NULL,
    block_id INTEGER NOT NULL,
    PRIMARY KEY (census_id, block_id),
    FOREIGN KEY (census_id) REFERENCES censuses(id) ON DELETE CASCADE,
    FOREIGN KEY (block_id) REFERENCES blocks(id) ON DELETE CASCADE
);
CREATE INDEX idx_census_blocks_census_id ON census_blocks(census_id);
CREATE INDEX idx_census_blocks_block_id ON census_blocks(block_id);

-- +goose Down
DROP INDEX IF EXISTS idx_census_blocks_block_id;
DROP INDEX IF EXISTS idx_census_blocks_census_id;
DROP INDEX IF EXISTS idx_strategy_tokens_token_id;
DROP INDEX IF EXISTS idx_strategy_tokens_strategy_id;
DROP INDEX IF EXISTS idx_token_holders_block_id;
DROP INDEX IF EXISTS idx_token_holders_holder_id;
DROP INDEX IF EXISTS idx_token_holders_token_id;
DROP INDEX IF EXISTS idx_censuses_strategy_id;
DROP INDEX IF EXISTS idx_tokens_type_id;

DROP TABLE IF EXISTS census_blocks;
DROP TABLE IF EXISTS strategy_tokens;
DROP TABLE IF EXISTS token_holders;
DROP TABLE IF EXISTS holders;
DROP TABLE IF EXISTS blocks;
DROP TABLE IF EXISTS censuses;
DROP TABLE IF EXISTS tokens;
DROP TABLE IF EXISTS token_types;
DROP TABLE IF EXISTS strategies;