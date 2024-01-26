
-- +goose Up

-- prepare token_holders table to delete blocks, holders and census_blocks tables
CREATE TABLE token_holders_backup (
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

INSERT INTO token_holders_backup SELECT * FROM token_holders;
DROP TABLE token_holders;
ALTER TABLE token_holders_backup RENAME TO token_holders;
DROP TABLE blocks;
DROP TABLE holders;
DROP TABLE census_blocks;

DELETE FROM token_types;
UPDATE sqlite_sequence SET seq = 0 WHERE name = 'token_types';
INSERT INTO token_types (id, type_name) VALUES (0, 'unknown');
INSERT INTO token_types (id, type_name) VALUES (1, 'erc20');
INSERT INTO token_types (id, type_name) VALUES (2, 'erc721');;
INSERT INTO token_types (id, type_name) VALUES (3, 'erc777');
INSERT INTO token_types (id, type_name) VALUES (4, 'poap');
INSERT INTO token_types (id, type_name) VALUES (5, 'gitcoinpassport');

DELETE FROM tokens WHERE type_id = 3;
DELETE FROM tokens WHERE type_id = 5;
DELETE FROM tokens WHERE type_id = 6;

UPDATE tokens SET type_id = 3 WHERE type_id = 4;
UPDATE tokens SET type_id = 4 WHERE type_id = 8;
UPDATE tokens SET type_id = 5 WHERE type_id = 100;

ALTER TABLE tokens ADD COLUMN last_block BIGINT NOT NULL DEFAULT 0;
ALTER TABLE tokens ADD COLUMN analysed_transfers BIGINT NOT NULL DEFAULT 0;

-- List of changes:
--  * Remove all token types and reset sequence
--    * Recreate token types with new ids:
--      * [new] unknown 0
--      * erc20 1
--      * erc721 2
--      * [updated] erc777 from 4 to 3
--      * [updated] poap from 8 to 4
--      * [updated] gitcoinpassport from 100 to 5
--      * [removed] erc1155
--      * [removed] nation3
--      * [removed] want
--      * [removed] erc721_burned
--  * Add 'last_block' column to 'tokens' table
--  * Add 'analysed_transfers' column to 'tokens' table