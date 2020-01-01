
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE user_banks (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE restrict,
    bank VARCHAR (100),
    branch VARCHAR (100),
    account_number VARCHAR (20),
    holdername VARCHAR (100),
    is_primary SMALLINT DEFAULT 0,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE user_banks;