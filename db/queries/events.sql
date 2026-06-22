-- name: GetEvent :one
SELECT * FROM events
WHERE id = ? LIMIT 1;

-- name: ListEvents :many
SELECT * FROM events
ORDER BY event_date DESC;

-- name: CreateEvent :one
INSERT INTO events (
  name, event_date, location
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateEvent :one
UPDATE events
SET name       = COALESCE(?, name),
    event_date = COALESCE(?, event_date),
    location   = COALESCE(?, location)
WHERE id = ?
RETURNING *;

-- name: DeleteEvent :execresult
DELETE FROM events
WHERE id = ?;