
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE poktan_stages (
    id SERIAL PRIMARY KEY,
    name VARCHAR (50),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Seeder

INSERT INTO poktan_stages (name) VALUES ('Daftar'), ('Follow Up'), ('Survey'), ('Terima'), ('Tolak');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE poktan_stages;
