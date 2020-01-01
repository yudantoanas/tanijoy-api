
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TYPE transaction_types as ENUM ('Send', 'Receive');

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    transaction_code VARCHAR (45),
    investor_id INTEGER NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
    related_id INTEGER,
    amount INTEGER,
    unique_code INTEGER,
    type transaction_types,
    note VARCHAR (191),
    payment_account VARCHAR (100),
    payment_method VARCHAR (100),
    transaction_at TIMESTAMP,
    confirmed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE transactions;

DROP TYPE transaction_types;
