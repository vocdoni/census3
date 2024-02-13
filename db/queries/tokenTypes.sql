-- name: ListTokenTypes :many
SELECT * FROM token_types
ORDER BY id;

-- name: CreateTokenType :execresult
INSERT INTO token_types (id, type_name) VALUES (?, ?);

-- name: UpdateTokenType :execresult
UPDATE token_types SET type_name = ? WHERE id = ?;