
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TYPE investment_update_categories AS enum ('Funded', 'Finished', 'Added');

CREATE TABLE investment_update(
    id SERIAL PRIMARY KEY,
    project_id INTEGER REFERENCES projects (id),
    related_id INTEGER,
    category investment_update_categories,
    additional_information VARCHAR (100),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE investment_update;

DROP TYPE investment_update_categories;