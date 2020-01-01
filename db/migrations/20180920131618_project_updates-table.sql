
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TYPE project_update_categories as ENUM ('Funded', 'Finished');

CREATE TABLE project_updates(
    id SERIAL PRIMARY KEY,
    project_id INTEGER,
    category project_update_categories,
    additional_information TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE project_updates;

DROP TYPE project_update_categories;
