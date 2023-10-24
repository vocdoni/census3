-- +goose Up

-- tokens table schema updates
ALTER TABLE tokens ADD COLUMN external_id TEXT;
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
    external_id TEXT NULL DEFAULT '',
    PRIMARY KEY (id, chain_id, external_id),
    FOREIGN KEY (type_id) REFERENCES token_types(id) ON DELETE CASCADE
);
INSERT INTO tokens_copy SELECT * FROM tokens;
DROP TABLE tokens;
-- DROP INDEX IF EXISTS idx_tokens_type_id;
ALTER TABLE tokens_copy RENAME TO tokens;
CREATE INDEX idx_tokens_type_id ON tokens(type_id);

-- token_holders table schema updates
ALTER TABLE token_holders ADD COLUMN external_id TEXT;
CREATE TABLE token_holders_copy (
    token_id BLOB NOT NULL,
    holder_id BLOB NOT NULL,
    balance BLOB NOT NULL,
    block_id INTEGER NOT NULL,
    external_id TEXT NULL DEFAULT '',
    PRIMARY KEY (token_id, holder_id, block_id, external_id),
    FOREIGN KEY (token_id) REFERENCES tokens(id) ON DELETE CASCADE,
    FOREIGN KEY (holder_id) REFERENCES holders(id) ON DELETE CASCADE,
    FOREIGN KEY (block_id) REFERENCES blocks(id) ON DELETE CASCADE
);
INSERT INTO token_holders_copy SELECT * FROM token_holders;
-- DROP INDEX IF EXISTS idx_token_holders_token_id;
-- DROP INDEX IF EXISTS idx_token_holders_holder_id;
-- DROP INDEX IF EXISTS idx_token_holders_block_id;
DROP TABLE token_holders;
ALTER TABLE token_holders_copy RENAME TO token_holders;
CREATE INDEX idx_token_holders_token_id ON token_holders(token_id);
CREATE INDEX idx_token_holders_holder_id ON token_holders(holder_id);
CREATE INDEX idx_token_holders_block_id ON token_holders(block_id);