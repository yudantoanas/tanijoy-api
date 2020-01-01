
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE investment_closing (
    id SERIAL PRIMARY KEY,
    investment_id INTEGER REFERENCES investments (id) ON DELETE CASCADE,
    statement TEXT,
    report_file VARCHAR (100),
    actual_return INTEGER,
    actual_percentage INTEGER,
    created_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE investment_closing;