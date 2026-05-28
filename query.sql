-- name: GetFighter :one
SELECT * FROM fighters
WHERE id = ? LIMIT 1;

-- name: ListFighters :many
SELECT * FROM fighters
ORDER BY name;

-- name: CreateFighter :one
INSERT INTO fighters (
  name, age, nickname
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateFighter :exec
UPDATE fighters
set name = ?,
age = ?,
nickname = ?
WHERE id = ?;

-- name: DeleteFighter :exec
DELETE FROM fighters
WHERE id = ?;

-- -- name: UpdateFighter :one
-- UPDATE fighters
-- set name = ?,
-- age = ?,
-- nickname = ?
-- WHERE id = ?
-- RETURNING *;