
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE farmer_commodity (
    farmer_id INTEGER REFERENCES farmers (id) ON DELETE CASCADE,
    commodity_id INTEGER REFERENCES commodities (id) ON DELETE CASCADE
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE farmer_commodity;
