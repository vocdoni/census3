-- +goose Up
CREATE TABLE latest_blocks (
    contract TEXT PRIMARY KEY,
    block_number BIGINT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS latest_blocks;