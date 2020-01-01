
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE job_position_categories(
    id SERIAL PRIMARY KEY,
    name VARCHAR (50),
    description varchar (191),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE job_position_categories;
