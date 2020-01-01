
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TYPE satisfaction_levels AS enum ('Sad', 'Neutral', 'Happy');

CREATE TABLE investment_review (
    id SERIAL PRIMARY KEY,
    investment_id INTEGER REFERENCES investments (id) ON DELETE CASCADE,
    satisfaction_level satisfaction_levels,
    review TEXT,
    is_shared SMALLINT,
    created_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE investment_review;

DROP TYPE satisfaction_levels;