CREATE TABLE IF NOT EXISTS fighters (
    id INTEGER PRIMARY KEY AUTO INCREMENT NOT NULL,
    name text NOT NULL,
    age INTEGER NOT NULL CHECK (age > 0 AND age <=255),
    nickname text
)