
-- +goose Up
DELETE FROM token_types;
UPDATE sqlite_sequence SET seq = 0 WHERE name = 'token_types';
INSERT INTO token_types (type_name) VALUES ('erc20');
INSERT INTO token_types (type_name) VALUES ('erc721');;
INSERT INTO token_types (type_name) VALUES ('erc777');
INSERT INTO token_types (type_name) VALUES ('poap');

ALTER TABLE tokens ADD COLUMN last_block BIGINT NOT NULL DEFAULT 0;
ALTER TABLE tokens ADD COLUMN analysed_transfers BIGINT NOT NULL DEFAULT 0;

-- List of changes:
--  * Remove all token types and reset sequence
--  * Add 'last_block' column to 'tokens' table
--  * Add 'analysed_transfers' column to 'tokens' table