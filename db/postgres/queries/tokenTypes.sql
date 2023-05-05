-- name: ListTokenTypes :many
SELECT * FROM token_types
ORDER BY id;

-- name: TokenTypeByID :one
SELECT * FROM token_types
WHERE id = $1
LIMIT 1;

-- name: TokenTypeByName :one
SELECT * FROM token_types
WHERE type_name = $1
LIMIT 1;

-- name: CreateTokenType :execresult
INSERT INTO token_types (type_name)
VALUES ($1);

-- name: UpdateTokenType :execresult
UPDATE token_types
SET type_name = sqlc.arg(type_name)
WHERE id = sqlc.arg(id);

-- name: DeleteTokenType :execresult
DELETE FROM token_types
WHERE id = $1;
