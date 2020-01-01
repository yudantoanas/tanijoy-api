
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
UPDATE proposals SET image_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/header/headercabaikeriting.jpg', thumbnail_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/card/cabai+keriting.jpg' WHERE id = 1;
UPDATE proposals SET image_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/header/headerbrokoli.jpg', thumbnail_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/card/brokoli.jpg' WHERE id = 2;
UPDATE proposals SET image_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/header/headerdaunbawang.jpg', thumbnail_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/card/daun+bawang.jpg' WHERE id = 3;
UPDATE proposals SET image_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/header/headerkacangkapri.jpg', thumbnail_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/card/kacang+kapri.jpg' WHERE id = 4;
UPDATE proposals SET image_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/header/headerpakchoy.jpg', thumbnail_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/card/pakchoy.jpg' WHERE id = 5;
UPDATE proposals SET image_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/header/headersawiputih.jpg', thumbnail_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/card/sawi+putih.jpg' WHERE id = 6;
UPDATE proposals SET image_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/header/headerkentang.jpg', thumbnail_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/card/potato.jpg' WHERE id = 8;
UPDATE proposals SET image_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/header/headertomat.jpg', thumbnail_path = 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/card/tomatbulat.jpg' WHERE id = 9;

DELETE FROM project_files WHERE id IN (7,8,9,10,11,12);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

