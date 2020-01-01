
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE bank_list (
    id SERIAL PRIMARY KEY,
    name VARCHAR (30),
    alias VARCHAR (20),
    logo_path VARCHAR (191),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

-- Seeder
INSERT INTO bank_list (name, alias, logo_path) VALUES
('Citibank', 'Citibank',''),
('Deutsche Bank AG', 'Deutsche Bank',''),
('Standard Chartered Bank', 'Standard Chartered',''),
('Bank HSBC Indonesia', 'HSBC',''),
('Bank ANZ Indonesia', 'ANZ',''),
('Bank Bukopin', 'Bukopin',''),
('Bank Central Asia', 'BCA','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/public/banks/bca.png'),
('Bank CIMB Niaga', 'CIMB','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/public/banks/cimb.png'),
('Bank CIMB Niaga Syariah', 'CIMB','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/public/banks/cimb-syariah.png'),
('Bank Comonwealth', 'Commonwealth',''),
('Bank Danamon Indonesia', 'Danamon',''),
('Bank DBS Indonesia', 'DBS',''),
('Bank ICBC Indonesia', 'ICBC',''),
('Bank Mandiri', 'Mandiri','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/public/banks/mandiri.png'),
('Bank Maybank Indonesia', 'Maybank',''),
('Bank Mayora', 'Mayora',''),
('Bank Mega', 'Mega',''),
('Bank Negara Indonesia', 'BNI','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/public/banks/bni.png'),
('Bank OCBC NISP', 'OCBC',''),
('Bank Panin Indonesia', 'Panin',''),
('Bank Permata', 'Permata',''),
('Bank Rakyat Indonesia', 'BRI','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/public/banks/bri.png'),
('Bank Tabungan Negara', 'BTN',''),
('Bank UOB Indonesia', 'UOB','');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE bank_list;
