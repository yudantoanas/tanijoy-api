
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE birthday_greet (
    investor_id INTEGER REFERENCES users (id) ON DELETE RESTRICT ,
    cs_id INTEGER REFERENCES users (id) ON DELETE RESTRICT,
    feedback TEXT NULL,
    call_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE birthday_greet;
