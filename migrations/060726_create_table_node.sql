-- +goose Up
CREATE TABLE IF NOT EXISTS nodes (
    node_id INTEGER PRIMARY KEY AUTOINCREMENT,
    label varchar(150) NOT NULL,
    x INTEGER NOT NULL,
    y INTEGER NOT NULL,
    created_timestamp TIMESTAMP NOT NULL,
    updated_timestamp TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE nodes;