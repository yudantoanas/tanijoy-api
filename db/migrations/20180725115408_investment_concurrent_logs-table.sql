
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TYPE log_activities as ENUM ('Invest', 'Pay', 'Verify', 'Release');

CREATE TABLE investment_logs (
    id SERIAL PRIMARY KEY,
    investor_id INTEGER REFERENCES users (id) ON DELETE SET NULL,
    activities log_activities,
    code VARCHAR (45),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE investment_logs;

DROP TYPE log_activities;