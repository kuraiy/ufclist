-- +goose Up
CREATE TABLE IF NOT EXISTS fighters (
    id INTEGER PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    age INTEGER NOT NULL,
    nickname TEXT
);
