-- +goose Up
CREATE TABLE users (
    fid INTEGER PRIMARY KEY, /* Farcaster ID */
    username TEXT NOT NULL UNIQUE,
    signer BLOB NOT NULL UNIQUE, /* Ed25519 public key */
    custodyAddress BLOB NOT NULL, /* EVM compatible address */
    appkeys BLOB, /* Keys which let apps write messages on the user behalf */
    recoveryAddress BLOB NOT NULL,
);
-- +goose Down
DROP TABLE IF EXISTS users;
