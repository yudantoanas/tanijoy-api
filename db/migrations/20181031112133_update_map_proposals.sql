
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO land_project (project_id, land_id, quantity, ready) VALUES
(7, 1, 4, 4),
(8, 1, 3, 3),
(9, 1, 3, 3);

INSERT INTO project_files (project_id, filename, path, label, purpose) VALUES
(1, 'maps-proyek', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/maps/Sukawangi.png', 'image', 'gallery'),
(2, 'maps-proyek', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/maps/Sukawangi.png', 'image', 'gallery'),
(3, 'maps-proyek', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/maps/Sukawangi.png', 'image', 'gallery'),
(4, 'maps-proyek', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/maps/Sukawangi.png', 'image', 'gallery'),
(5, 'maps-proyek', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/maps/Sukawangi.png', 'image', 'gallery'),
(6, 'maps-proyek', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/maps/Sukawangi.png', 'image', 'gallery'),
(7, 'maps-proyek', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/maps/Sukawangi.png', 'image', 'gallery'),
(8, 'maps-proyek', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/maps/Berastagi.png', 'image', 'gallery'),
(9, 'maps-proyek', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/maps/Sukawangi.png', 'image', 'gallery');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

