-- +goose Up
CREATE TABLE token_updates (
    id BLOB NOT NULL,
    chain_id INTEGER NOT NULL DEFAULT 0,
    filter_gob BLOB NOT NULL DEFAULT '',
    last_block BIGINT NOT NULL DEFAULT 0,
    PRIMARY KEY (id, chain_id, external_id),
);