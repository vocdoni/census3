-- name: PaginatedTokenTypes :many
SELECT * FROM TokenTypes
ORDER BY id
LIMIT ? OFFSET ?;

-- name: TokenTypeByID :one
SELECT * FROM TokenTypes
WHERE id = ?
LIMIT 1;

-- name: TokenTypeByName :one
SELECT * FROM TokenTypes
WHERE type_name = ?
LIMIT 1;

-- name: CreateTokenType :execresult
INSERT INTO TokenTypes (type_name)
VALUES (?);

-- name: UpdateTokenType :execresult
UPDATE TokenTypes
SET type_name = sqlc.arg(type_name)
WHERE id = sqlc.arg(id);

-- name: DeleteTokenType :execresult
DELETE FROM TokenTypes
WHERE id = ?;
