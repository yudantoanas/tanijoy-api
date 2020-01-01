
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
UPDATE commodities set name = 'Cabai Keriting', type = 'Cabai' where id = 1;

INSERT INTO commodities (name, type) VALUES
('Cabai Rawit', 'Cabai'),
('Tomat', 'Tomat');

INSERT INTO proposals (commodities_id, name, type, amount, return_expected, return_min, return_max, area, period, plants_count, risk, image_path, is_featured) VALUES
(8, 'Cabai Rawit Bogor', 'Budidaya', 22275000, 18, 5, 25, 2500, 180, 6000, 'Tinggi', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-cabai-rawit-rectangle.jpg',0),
(7, 'Kentang Granola Berastagi', 'Budidaya', 22275000, 12, 5, 25, 1000, 180, 2500, 'Tinggi', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-kentang-granola-rectangle.jpg',0),
(9, 'Tomat Bogor', 'Budidaya', 22275000, 12 , 5, 25, 500, 180, 2500, 'Tinggi', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-tomat-buah-rectangle.jpg',0);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

