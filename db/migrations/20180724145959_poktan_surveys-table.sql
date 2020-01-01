
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE poktan_surveys(
    id SERIAL PRIMARY KEY,
    candidate_id INTEGER REFERENCES poktan_candidates (id) ON DELETE SET NULL,
    surveyor_id INTEGER REFERENCES users (id) ON DELETE SET NULL,
    appointment_at TIMESTAMP,
    note TEXT,
    next VARCHAR (191),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE poktan_surveys;
