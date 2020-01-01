
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users ALTER COLUMN phone DROP NOT NULL;
ALTER TABLE users ALTER COLUMN username DROP NOT NULL;

INSERT INTO users (name, email, password, confirmed, confirmation_code) VALUES
('Lorem Tanijoy', 'lorem@tanijoy.id', 'tumbuhbahagia', 1, 'loremipsumdolorsitamet');

UPDATE users SET phone = '', username = '', password = 'tumbuhbahagia', confirmation_code = 'tumbuhbahagia'
WHERE email = 'alfian@tanijoy.id';

INSERT INTO role_user (user_id, role_id) VALUES
(1,1);

INSERT INTO user_information (user_id, identities_number) VALUES
(1, '3515080902950002');

INSERT INTO user_files (user_id, purpose, path) VALUES
(1, 'Identities', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/public/users/image/ktp_dummy.jpg');

INSERT INTO user_banks (user_id, bank, branch, account_number, holdername) VALUES
(1, 'BCA', 'Jakarta', '11829123', 'Lorem Tanijoy');

INSERT INTO farmers (account_id, name, phone, address, age) VALUES
(null, 'Majid', null, 'Kp. Cetang Malang, ds Sukawangi, kec. Sukamakmur, Kab. Bogor', 55);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back