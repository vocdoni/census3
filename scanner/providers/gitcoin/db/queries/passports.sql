-- name: NewScore :execresult
INSERT INTO scores (address, score, date) VALUES (?, ?, ?);

-- name: GetScores :many
SELECT address, score FROM scores;

-- name: GetScore :one
SELECT score FROM scores WHERE address = sqlc.arg(address);

-- name: ScoreExists :one
SELECT EXISTS (
    SELECT address FROM scores WHERE address = sqlc.arg(address)
);

-- name: UpdateScore :execresult
UPDATE scores 
SET score = sqlc.arg(score),
    date = sqlc.arg(date)
WHERE address = sqlc.arg(address);

-- name: DeleteScore :execresult
DELETE FROM scores WHERE address = sqlc.arg(address);

-- name: NewStampScore :execresult
INSERT INTO stamps (address, name, score) VALUES (?, ?, ?);

-- name: GetStampScores :many
SELECT address, score FROM stamps WHERE name = sqlc.arg(name);

-- name: StampScoreExists :one
SELECT EXISTS (
    SELECT address FROM stamps 
    WHERE address = sqlc.arg(address) 
        AND name = sqlc.arg(name)
);

-- name: UpdateStampScore :execresult
UPDATE stamps
SET score = sqlc.arg(score)
WHERE address = sqlc.arg(address) 
    AND name = sqlc.arg(name);

-- name: GetStampsForAddress :many
SELECT name, score FROM stamps WHERE address = sqlc.arg(address);

-- name: DeleteStampForAddress :execresult
DELETE FROM stamps WHERE address = sqlc.arg(address);