-- name: ListTokenTypes :many
SELECT * FROM token_types
ORDER BY id;

-- name: TokenTypeByID :one
SELECT * FROM token_types
WHERE id = ?
LIMIT 1;

-- name: TokenTypeByName :one
SELECT * FROM token_types
WHERE type_name = ?
LIMIT 1;

-- name: CreateTokenType :execresult
INSERT INTO token_types (type_name)
VALUES (?);

-- name: UpdateTokenType :execresult
UPDATE token_types
SET type_name = sqlc.arg(type_name)
WHERE id = sqlc.arg(id);

-- name: DeleteTokenType :execresult
DELETE FROM token_types
WHERE id = ?;
