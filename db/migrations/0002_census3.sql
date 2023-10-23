-- +goose Up

-- tokens table schema updates
ALTER TABLE tokens ADD COLUMN meta_token_id TEXT;
CREATE TABLE tokens_copy (
    id BLOB PRIMARY KEY NOT NULL,
    name TEXT,
    symbol TEXT,
    decimals INTEGER,
    total_supply BLOB,
    creation_block BIGINT,
    type_id INTEGER NOT NULL,
    synced BOOLEAN NOT NULL,
    tags TEXT,
    chain_id INTEGER NOT NULL,
    meta_token_id BLOB NULL DEFAULT '',
    PRIMARY KEY (id, chain_id),
    FOREIGN KEY (type_id) REFERENCES token_types(id) ON DELETE CASCADE
);
INSERT INTO tokens_copy SELECT * FROM tokens;
DROP TABLE tokens;
-- DROP INDEX IF EXISTS idx_tokens_type_id;
ALTER TABLE tokens_copy RENAME TO tokens;
CREATE INDEX idx_tokens_type_id ON tokens(type_id);
