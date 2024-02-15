-- name: CreateFidAppKey :execresult
INSERT INTO fid_appkeys (fid, app_key) VALUES (?, ?);

-- name: DeleteFidAppKey :execresult
DELETE FROM fid_appkeys WHERE fid = ? AND app_key = ?;

-- name: GetFidAppKeys :many
SELECT app_key FROM fid_appkeys WHERE fid = ?;

-- name: GetFidByAppKey :many
SELECT fid FROM fid_appkeys WHERE app_key = ?;
