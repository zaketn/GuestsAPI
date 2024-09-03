CREATE TABLE IF NOT EXISTS guests
(
    id        serial PRIMARY KEY,
    name      VARCHAR(128),
    last_name VARCHAR(128),
    email     VARCHAR(128) UNIQUE,
    phone     VARCHAR(128) UNIQUE,
    country   VARCHAR(128)
)