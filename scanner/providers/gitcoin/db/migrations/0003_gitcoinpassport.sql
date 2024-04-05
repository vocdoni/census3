-- +goose Up
CREATE TABLE total_supplies (
    name TEXT PRIMARY KEY,
    total_supply TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS total_supplies;