-- +goose Up
CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    event_date DATE NOT NULL,
    location TEXT
);