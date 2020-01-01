
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE poktan_candidates ALTER COLUMN area TYPE INTEGER USING (trim(area)::INTEGER );

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

