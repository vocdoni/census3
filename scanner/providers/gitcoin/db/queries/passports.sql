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

-- name: TotalSupplyScores :many
SELECT score FROM scores;

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

-- name: GetStampScoreForAddress :one
SELECT score 
FROM stamps 
WHERE address = sqlc.arg(address) 
    AND name = sqlc.arg(stamp);

-- name: DeleteStampForAddress :execresult
DELETE FROM stamps WHERE address = sqlc.arg(address);

-- name: AvailableStamps :many
SELECT DISTINCT(name) FROM stamps;

-- name: ExistsStamp :one
SELECT EXISTS (
    SELECT name FROM (
        SELECT DISTINCT(name) FROM stamps
    ) WHERE name = sqlc.arg(stamp)
);

-- name: StampTotalSupplyScores :many
SELECT score FROM stamps WHERE name = sqlc.arg(stamp);