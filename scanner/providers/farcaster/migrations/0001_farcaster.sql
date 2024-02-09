-- +goose Up
CREATE TABLE users (
    fid INTEGER PRIMARY KEY, /* Farcaster ID */
    username TEXT NOT NULL DEFAULT '',
    signer BLOB, /* Ed25519 public key */
    custody_address BLOB NOT NULL, /* EVM compatible address */
    app_keys BLOB, /* Keys which let apps write messages on the user behalf */
    recovery_address BLOB NOT NULL
);
-- +goose Down
DROP TABLE IF EXISTS users;
