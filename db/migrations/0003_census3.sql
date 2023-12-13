-- +goose Up

-- add created_at column to tokens table
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
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id, chain_id, external_id),
    FOREIGN KEY (type_id) REFERENCES token_types(id) ON DELETE CASCADE,
    FOREIGN KEY (default_strategy) REFERENCES strategies(id) ON DELETE CASCADE
);
INSERT INTO tokens_copy (id, name, symbol, decimals, total_supply, creation_block, type_id, synced, tags, chain_id, chain_address, external_id, default_strategy, icon_uri)
SELECT * FROM tokens;
DROP TABLE tokens;
ALTER TABLE tokens_copy RENAME TO tokens;
CREATE INDEX idx_tokens_type_id ON tokens(type_id);