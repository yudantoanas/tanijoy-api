
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TYPE gender_choices AS enum ('M', 'F');

CREATE TABLE user_information(
    user_id INTEGER REFERENCES users (id) ON DELETE restrict,
    identities_number VARCHAR (20),
    province VARCHAR (100),
    city VARCHAR (100),
    address TEXT,
    gender gender_choices,
    postal_code INTEGER,
    profile_progress INTEGER NULL,
    birth_at DATE,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user_information;

DROP TYPE gender_choices;
