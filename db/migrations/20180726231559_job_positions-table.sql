
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TYPE job_position_labels as ENUM ('Internship', 'Career');

CREATE TABLE job_positions(
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES job_position_categories (id) ON DELETE SET NULL,
    name VARCHAR (50),
    description TEXT,
    requirements TEXT,
    label job_position_labels,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE job_positions;

DROP TYPE job_position_labels;