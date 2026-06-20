-- +goose Up
CREATE TABLE IF NOT EXISTS maps (
                       map_id INTEGER PRIMARY KEY AUTOINCREMENT,
                       map_path varchar(100) NOT NULL UNIQUE,
                       created_timestamp TIMESTAMP NOT NULL,
                       updated_timestamp TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE maps;