
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TYPE survey_labels as ENUM ('document', 'photo', 'video');

CREATE TABLE poktan_survey_files (
    id SERIAL PRIMARY KEY,
    survey_id INTEGER REFERENCES poktan_surveys (id) ON DELETE restrict,
    candidate_id INTEGER REFERENCES poktan_candidates (id) ON DELETE SET NULL,
    filename VARCHAR (191),
    label survey_labels,
    path VARCHAR (191),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE poktan_survey_files;

DROP TYPE survey_labels;