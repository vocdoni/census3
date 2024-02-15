-- +goose Up
CREATE TABLE users (
    fid INTEGER PRIMARY KEY
);

CREATE TABLE fid_appkeys (
    fid INTEGER NOT NULL,
    app_key TEXT NOT NULL,
    PRIMARY KEY (fid, app_key),
    FOREIGN KEY (fid) REFERENCES users(fid) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS fid_appkeys;

