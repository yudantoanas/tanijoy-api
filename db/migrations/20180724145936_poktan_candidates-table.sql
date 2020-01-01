
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE poktan_candidates (
    id SERIAL PRIMARY KEY,
    email VARCHAR (100),
    name VARCHAR (100),
    phone VARCHAR (20),
    address TEXT,
    area VARCHAR (20),
    farmers_count INTEGER,
    vegetables VARCHAR (100),
    stage_id INTEGER REFERENCES poktan_stages (id) ON DELETE set NULL,
    contacted_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE poktan_candidates;
