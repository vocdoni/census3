-- name: PaginatedTokenTypes :many
SELECT * FROM TokenTypes
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: TokenTypeByID :one
SELECT * FROM TokenTypes
WHERE id = $1
LIMIT 1;

-- name: TokenTypeByName :one
SELECT * FROM TokenTypes
WHERE type_name = $1
LIMIT 1;

-- name: CreateTokenType :execresult
INSERT INTO TokenTypes (type_name)
VALUES ($1);

-- name: UpdateTokenType :execresult
UPDATE TokenTypes
SET type_name = sqlc.arg(type_name)
WHERE id = sqlc.arg(id);

-- name: DeleteTokenType :execresult
DELETE FROM TokenTypes
WHERE id = $1;
