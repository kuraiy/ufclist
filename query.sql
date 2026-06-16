-- name: GetFighter :one
SELECT * FROM fighters
WHERE id = ? LIMIT 1;

-- name: ListFighters :many
SELECT * FROM fighters  
WHERE activated = 1 
ORDER BY name;

-- name: CreateFighter :one
INSERT INTO fighters (
  name, age, nickname
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateFighter :execresult
UPDATE fighters
set name = ?,
age = ?,
nickname = ?
WHERE id = ?;

-- name: DeleteFighter :execresult
UPDATE fighters
set activated = 0
WHERE id = ?;
