-- +goose Up
CREATE TABLE strategies (
    id BIGSERIAL PRIMARY KEY,
    predicate TEXT NOT NULL
);

CREATE TABLE token_types (
    id BIGSERIAL PRIMARY KEY,
    type_name TEXT NOT NULL UNIQUE
);

INSERT INTO token_types (type_name) VALUES ('erc20');
INSERT INTO token_types (type_name) VALUES ('erc721');
INSERT INTO token_types (type_name) VALUES ('erc1155');
INSERT INTO token_types (type_name) VALUES ('erc777');
INSERT INTO token_types (type_name) VALUES ('nation3');
INSERT INTO token_types (type_name) VALUES ('want');

CREATE TABLE tokens (
    id BYTEA PRIMARY KEY NOT NULL,
    name TEXT,
    symbol TEXT,
    decimals BIGINT,
    total_supply BYTEA,
    creation_block BIGINT NOT NULL,
    type_id BIGINT NOT NULL,
    FOREIGN KEY (type_id) REFERENCES token_types(id) ON DELETE CASCADE
);
CREATE INDEX idx_tokens_type_id ON tokens(type_id);

CREATE TABLE censuses (
    id BIGSERIAL PRIMARY KEY,
    strategy_id BIGINT NOT NULL,
    merkle_root BYTEA NOT NULL UNIQUE,
    uri TEXT UNIQUE,
    FOREIGN KEY (strategy_id) REFERENCES strategies(id) ON DELETE CASCADE
);
CREATE INDEX idx_censuses_strategy_id ON censuses(strategy_id);

CREATE TABLE blocks (
    id BIGINT PRIMARY KEY NOT NULL,
    timestamp TEXT NOT NULL UNIQUE,
    root_hash BYTEA NOT NULL UNIQUE
);

CREATE TABLE holders (
    id BYTEA PRIMARY KEY NOT NULL
);

CREATE TABLE token_holders (
    token_id BYTEA NOT NULL,
    holder_id BYTEA NOT NULL,
    balance BYTEA NOT NULL,
    block_id BIGINT NOT NULL,
    PRIMARY KEY (token_id, holder_id, block_id),
    FOREIGN KEY (token_id) REFERENCES tokens(id) ON DELETE CASCADE,
    FOREIGN KEY (holder_id) REFERENCES holders(id) ON DELETE CASCADE,
    FOREIGN KEY (block_id) REFERENCES blocks(id) ON DELETE CASCADE
);
CREATE INDEX idx_token_holders_token_id ON token_holders(token_id);
CREATE INDEX idx_token_holders_holder_id ON token_holders(holder_id);
CREATE INDEX idx_token_holders_block_id ON token_holders(block_id);

CREATE TABLE strategy_tokens (
    strategy_id BIGINT NOT NULL,
    token_id BYTEA NOT NULL,
    min_balance BYTEA NOT NULL,
    method_hash BYTEA NOT NULL,
    PRIMARY KEY (strategy_id, token_id),
    FOREIGN KEY (strategy_id) REFERENCES strategies(id) ON DELETE CASCADE,
    FOREIGN KEY (token_id) REFERENCES tokens(id) ON DELETE CASCADE
);
CREATE INDEX idx_strategy_tokens_strategy_id ON strategy_tokens(strategy_id);
CREATE INDEX idx_strategy_tokens_token_id ON strategy_tokens(token_id);

CREATE TABLE census_blocks (
    census_id BIGINT NOT NULL,
    block_id BIGINT NOT NULL,
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
