
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
UPDATE proposals SET image_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-cabai-kriting-rectangle.jpg', thumbnail_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-cabai-kriting-rectangle.jpg' WHERE id = 1;
UPDATE proposals SET image_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-brokoli-rectangle.jpg', thumbnail_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-brokoli-rectangle.jpg' WHERE id = 2;
UPDATE proposals SET image_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-daun-bawang-rectangle.jpg', thumbnail_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-daun-bawang-rectangle.jpg' WHERE id = 3;
UPDATE proposals SET image_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-kacang-kapri-rectangle.jpg', thumbnail_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-kacang-kapri-rectangle.jpg' WHERE id = 4;
UPDATE proposals SET image_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-pak-choy-rectangle.jpg', thumbnail_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-pak-choy-rectangle.jpg' WHERE id = 5;
UPDATE proposals SET image_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-sawi-putih-rectangle.jpg', thumbnail_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-sawi-putih-rectangle.jpg' WHERE id = 6;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

