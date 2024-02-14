-- +goose Up
CREATE TABLE users (
    fid INTEGER PRIMARY KEY, /* Farcaster ID */
    signer BLOB NOT NULL, /* ED25519 public key */
    custody_address BLOB NOT NULL, /* EVM compatible address */
    app_keys BLOB, /* Keys which let apps write messages on the user behalf */
    recovery_address BLOB NOT NULL /* EVM compatible address */
);

CREATE TABLE linkedevm_fid (
    fid INTEGER NOT NULL,
    evm_address BLOB NOT NULL,
    PRIMARY KEY (fid, evm_address),
    FOREIGN KEY (fid) REFERENCES users(fid) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS linkedevm_fid;
