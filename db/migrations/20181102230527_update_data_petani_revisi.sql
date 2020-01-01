
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
TRUNCATE TABLE farmer_commodity RESTART IDENTITY;

DELETE FROM proposals WHERE id = 7;

INSERT INTO farmers (farmer_type, name, year_of_experiences, specializations, profpic_src) VALUES
('Petani Utama', 'Pak Parna Djawak', 5, 'Kentang Granola,Kubis,Cabai,Bawang,Jeruk','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/kentang-granola/Berastagi+-+Parna+Djawak.jpeg'),
('Petani Utama', 'Mang Oman', 10, 'Tomat,Brokoli,Cabai Keriting','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/tomat-bogor/Mang+Oman_+Tomat%2C++Brokoli%2C+Cabai+Keriting_10+Tahun.jpg');

UPDATE farmers SET specializations = 'Brokoli,Tomat,Daun Bawang' WHERE id = 2;
UPDATE farmers SET specializations = 'Kentang Granola,Cabai,Ubi Jepang,Bawang' WHERE id = 19;
UPDATE farmers SET specializations = 'Kentang Granola,Tomat,Kubis,Jeruk,Cabai,Bawang' WHERE id = 20;
UPDATE farmers SET specializations = 'Kentang Granola,Tomat,Cabai,Jeruk' WHERE id = 22;
UPDATE farmers SET specializations = 'Kentang Granola,Lobak,Cabai,Kubis', year_of_experiences = 25 WHERE id = 21;
UPDATE farmers SET specializations = 'Tomat,Sawi Putih,Kol' WHERE id = 28;
UPDATE farmers SET specializations = 'Tomat,Sawi Putih,Brokoli' WHERE id = 29;

INSERT INTO farmer_commodity (farmer_id, commodity_id) VALUES
(2,1),(4,1),(6,1),(7,1),(8,1),(9,1),(10,1),(11,1),
(12,2),(13,2),(14,2),
(1,3),(2,3),(3,3),(4,3),(5,3),
(15,4),(16,4),(17,4),(18,4),
(23,5),(24,5),
(24,6),(25,6),(26,6),
(19,7),(20,7),(21,7),(22,7),(30,7),
(15,9),(1,9),(7,9),(17,9),(28,9),(2,9),(25,9),(31,9),(29,9),(24,9);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

