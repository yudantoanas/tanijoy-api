
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TYPE file_purposes AS enum ('Identities', 'Profpic', 'Other');

CREATE TYPE file_types AS enum ('Image', 'Document');

CREATE TABLE user_files(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id) ON DELETE restrict,
    purpose file_purposes,
    type file_types,
    path VARCHAR (191),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user_files;

DROP TYPE file_purposes;

DROP TYPE file_types;