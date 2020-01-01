
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE farmer_land(
    farmer_id INTEGER REFERENCES farmers (id) ON DELETE CASCADE,
    land_id INTEGER REFERENCES lands (id) ON DELETE CASCADE
);

-- Seeder
TRUNCATE TABLE farmer_land;

INSERT INTO farmer_land VALUES
(1, 1), (2, 1), (3, 1), (4, 1),
(5, 1), (6, 1), (7, 1),
(8, 1);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE farmer_land;