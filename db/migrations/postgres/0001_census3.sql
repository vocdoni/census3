-- +goose Up
CREATE TABLE Strategies (
    id BIGINT PRIMARY KEY,
    predicate TEXT NOT NULL
);

CREATE TABLE TokenTypes (
    id BIGINT PRIMARY KEY,
    type_name TEXT NOT NULL UNIQUE
);

CREATE TABLE Tokens (
    id BYTEA PRIMARY KEY NOT NULL,
    name TEXT,
    symbol TEXT,
    decimals BIGINT,
    total_supply BYTEA,
    creation_block BIGINT NOT NULL,
    type_id BIGINT NOT NULL,
    FOREIGN KEY (type_id) REFERENCES TokenTypes(id) ON DELETE CASCADE
);
CREATE INDEX idx_tokens_type_id ON Tokens(type_id);

CREATE TABLE Censuses (
    id BIGINT PRIMARY KEY,
    strategy_id BIGINT NOT NULL,
    merkle_root BYTEA NOT NULL UNIQUE,
    uri TEXT UNIQUE,
    FOREIGN KEY (strategy_id) REFERENCES Strategies(id) ON DELETE CASCADE
);
CREATE INDEX idx_censuses_strategy_id ON Censuses(strategy_id);

CREATE TABLE Blocks (
    id BIGINT PRIMARY KEY NOT NULL,
    timestamp TEXT NOT NULL UNIQUE,
    root_hash BYTEA NOT NULL UNIQUE
);

CREATE TABLE Holders (
    id BYTEA PRIMARY KEY NOT NULL
);

CREATE TABLE TokenHolders (
    token_id BYTEA NOT NULL,
    holder_id BYTEA NOT NULL,
    balance BYTEA NOT NULL,
    block_id BIGINT NOT NULL,
    PRIMARY KEY (token_id, holder_id, block_id),
    FOREIGN KEY (token_id) REFERENCES Tokens(id) ON DELETE CASCADE,
    FOREIGN KEY (holder_id) REFERENCES Holders(id) ON DELETE CASCADE,
    FOREIGN KEY (block_id) REFERENCES Blocks(id) ON DELETE CASCADE
);
CREATE INDEX idx_tokenholders_token_id ON TokenHolders(token_id);
CREATE INDEX idx_tokenholders_holder_id ON TokenHolders(holder_id);
CREATE INDEX idx_tokenholders_block_id ON TokenHolders(block_id);

CREATE TABLE StrategyTokens (
    strategy_id BIGINT NOT NULL,
    token_id BYTEA NOT NULL,
    min_balance BYTEA NOT NULL,
    method_hash BYTEA NOT NULL,
    PRIMARY KEY (strategy_id, token_id),
    FOREIGN KEY (strategy_id) REFERENCES Strategies(id) ON DELETE CASCADE,
    FOREIGN KEY (token_id) REFERENCES Tokens(id) ON DELETE CASCADE
);
CREATE INDEX idx_strategytokens_strategy_id ON StrategyTokens(strategy_id);
CREATE INDEX idx_strategytokens_token_id ON StrategyTokens(token_id);

CREATE TABLE CensusBlocks (
    census_id BIGINT NOT NULL,
    block_id BIGINT NOT NULL,
    PRIMARY KEY (census_id, block_id),
    FOREIGN KEY (census_id) REFERENCES Censuses(id) ON DELETE CASCADE,
    FOREIGN KEY (block_id) REFERENCES Blocks(id) ON DELETE CASCADE
);
CREATE INDEX idx_censusblocks_census_id ON CensusBlocks(census_id);
CREATE INDEX idx_censusblocks_block_id ON CensusBlocks(block_id);

-- +goose Down
DROP INDEX IF EXISTS idx_censusblocks_block_id;
DROP INDEX IF EXISTS idx_censusblocks_census_id;
DROP INDEX IF EXISTS idx_strategytokens_token_id;
DROP INDEX IF EXISTS idx_strategytokens_strategy_id;
DROP INDEX IF EXISTS idx_tokenholders_block_id;
DROP INDEX IF EXISTS idx_tokenholders_holder_id;
DROP INDEX IF EXISTS idx_tokenholders_token_id;
DROP INDEX IF EXISTS idx_censuses_strategy_id;
DROP INDEX IF EXISTS idx_tokens_type_id;

DROP TABLE IF EXISTS CensusBlocks;
DROP TABLE IF EXISTS StrategyTokens;
DROP TABLE IF EXISTS TokenHolders;
DROP TABLE IF EXISTS Holders;
DROP TABLE IF EXISTS Blocks;
DROP TABLE IF EXISTS Censuses;
DROP TABLE IF EXISTS Tokens;
DROP TABLE IF EXISTS TokenTypes;
DROP TABLE IF EXISTS Strategies;
