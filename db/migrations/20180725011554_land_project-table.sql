
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE land_project (
    id SERIAL PRIMARY KEY,
    project_id INTEGER NOT NULL REFERENCES projects (id) ON DELETE CASCADE,
    land_id INTEGER REFERENCES lands (id) ON DELETE CASCADE,
    quantity INTEGER,
    hold INTEGER DEFAULT 0,
    ready INTEGER CHECK (ready >= 0),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    CHECK (quantity >= ready)
);

INSERT INTO land_project (project_id, land_id, quantity, ready) VALUES
(1, 1, 4, 4),
(1, 2, 3, 3),
(2, 1, 3, 3),
(2, 2, 2, 2),
(3, 1, 9, 9),
(3, 2, 5, 5),
(4, 1, 3, 3),
(5, 1, 3, 3),
(6, 1, 3, 3);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE land_project;
