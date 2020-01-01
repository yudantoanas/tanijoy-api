
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE regionals (
    id SERIAL PRIMARY KEY,
    manager_id INTEGER REFERENCES users (id) ON DELETE SET NULL,
    name VARCHAR (100),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

-- seeder
TRUNCATE TABLE regionals;

INSERT INTO regionals (manager_id, name) VALUES 
(7, 'Jawa Barat'), (8, 'Medan');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE regionals;
