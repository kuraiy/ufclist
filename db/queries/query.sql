-- name: GetFighter :one
SELECT id, name, age, nickname FROM fighters
WHERE id = ? AND activated = 1
LIMIT 1;

-- name: ListFighters :many
SELECT id, name, age, nickname FROM fighters
WHERE activated = 1
ORDER BY name;

-- name: CreateFighter :one
INSERT INTO fighters (
  name, age, nickname
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateFighter :one
UPDATE fighters
SET name     = COALESCE(?, name),
    age      = COALESCE(?, age),
    nickname = COALESCE(?, nickname)
WHERE id = ?
RETURNING id, name, age, nickname;

-- name: DeleteFighter :execresult
UPDATE fighters
set activated = 0
WHERE id = ?;
