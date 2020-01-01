
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE commodities(
    id SERIAL PRIMARY KEY,
    name VARCHAR (191) ,
    type VARCHAR (191),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

-- seeder
TRUNCATE TABLE commodities;

INSERT INTO commodities (name, type) VALUES
('Cabe Keriting', 'Cabe'),
('Daun Bawang', 'Bawang'),
('Brokoli', 'Brokoli'),
('Kacang Kapri', 'Kacang'),
('Pak Choy', 'Kubis'),
('Sawi Putih', 'Kubis'),
('Kentang', 'Kentang');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE commodities;
