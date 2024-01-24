
-- +goose Up
ALTER TABLE censuses ADD COLUMN accuracy FLOAT NOT NULL DEFAULT '100.0';


-- List of changes:
--  * Add 'accuracy' column to 'censuses' table