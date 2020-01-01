
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE investment_call_reports (
    id SERIAL PRIMARY KEY,
    investment_id INTEGER REFERENCES investments (id) ON DELETE SET NULL,
    investment_activity_id INTEGER REFERENCES investment_activities (id) ON DELETE SET NULL,
    call_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE investment_call_reports;