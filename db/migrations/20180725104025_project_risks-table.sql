
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE project_risks (
    id SERIAL PRIMARY KEY,
    project_id INTEGER NOT NULL REFERENCES projects (id) ON DELETE CASCADE,
    risk VARCHAR (50),
    mitigation VARCHAR (191),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE project_risks;
