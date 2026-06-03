-- +goose Up
CREATE TABLE IF NOT EXISTS fighters (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name text NOT NULL,
    age INTEGER NOT NULL,
    nickname text
);

-- +goose Down
DROP TABLE fighters;