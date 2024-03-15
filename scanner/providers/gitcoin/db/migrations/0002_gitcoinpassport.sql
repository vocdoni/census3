-- +goose Up
CREATE TABLE metadata (
    attr TEXT PRIMARY KEY,
    value TEXT NOT NULL
);

-- +goose Down
DROP INDEX IF EXISTS idx_metadata_key;
DROP TABLE IF EXISTS metadata;