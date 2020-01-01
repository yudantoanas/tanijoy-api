
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE land_concurrents (
    id SERIAL PRIMARY KEY,
    concurrent_id INTEGER NOT NULL REFERENCES investment_concurrents (id) ON DELETE CASCADE,
    land_id INTEGER NOT NULL REFERENCES lands (id) ON DELETE SET NULL,
    project_id INTEGER NOT NULL REFERENCES projects (id) on DELETE CASCADE,
    quantity INTEGER
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE land_concurrents;
