-- +goose Up
CREATE TABLE scores (
    address TEXT PRIMARY KEY,
    score TEXT NOT NULL,
    date TIMESTAMP NOT NULL
);

CREATE TABLE stamps (
    address TEXT NOT NULL,
    name TEXT NOT NULL,
    score TEXT NOT NULL,
    PRIMARY KEY (address, name),
    FOREIGN KEY (address) REFERENCES scores(address) ON DELETE CASCADE
);
CREATE INDEX idx_stamps_address ON stamps(address);
CREATE INDEX idx_stamps_name ON stamps(name);

-- +goose Down
DROP INDEX IF EXISTS idx_stamps_address;
DROP INDEX IF EXISTS idx_stamps_name;
DROP TABLE IF EXISTS stamps;
DROP TABLE IF EXISTS scores;