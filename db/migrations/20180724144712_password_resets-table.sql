
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE password_resets (
    email VARCHAR (100),
    token varchar (191),
    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE password_resets;
