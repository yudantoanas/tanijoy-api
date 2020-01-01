
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE lands 
    ADD contact_person VARCHAR(100),
    ADD contact_number VARCHAR(15);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE lands
    DROP COLUMN contact_person,
    DROP COLUMN contact_number;