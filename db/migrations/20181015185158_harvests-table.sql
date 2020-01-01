
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE harvests (
    id SERIAL PRIMARY KEY,
    investment_id INTEGER REFERENCES investments (id) ON DELETE SET NULL,
    qty INTEGER,
    price INTEGER
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE harvests;