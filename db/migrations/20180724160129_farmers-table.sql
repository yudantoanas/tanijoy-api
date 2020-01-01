
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TYPE farmer_types as ENUM ('Petani Utama', 'PHL');
CREATE TYPE marital_statuses as ENUM ('Belum Menikah', 'Menikah', 'Menikah (Duda)', 'Menikah (Janda)');

CREATE TABLE farmers (
    id SERIAL PRIMARY KEY,
    account_id INTEGER REFERENCES users (id) ON DELETE RESTRICT,
    farmer_type farmer_types,
    name VARCHAR (100),
    ktp VARCHAR (20),
    age INTEGER,
    phone VARCHAR (20),
    address TEXT,
    year_of_experiences INTEGER,
    marital_status marital_statuses,
    number_of_childs INTEGER,
    specializations VARCHAR (191),
    profpic_src TEXT,
    created_at TIMESTAMP DEFAULT current_timestamp,
    udpated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

-- Seeder
TRUNCATE TABLE farmers;

INSERT INTO farmers (farmer_type, name, ktp, age, phone, marital_status, number_of_childs, address, specializations, year_of_experiences, profpic_src) VALUES
('Petani Utama', 'Mang Yadi', '3201090405880008', NULL, '085714720082',  'Menikah', 2, 'Kampung Cikuray, RT 02 RW 05, Desa Sukawangi, Kec. Sukamakmur, Kab. Bogor, Kab. Bogor', 'Cabai Keriting, Tomat', NULL, 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/Bogor+-+Mang+Yadi.jpg'),
('Petani Utama', 'Mang Eman', '3201091207910006', 27, NULL,  'Belum Menikah', 0, 'Kampung Cikuray, RT 02 RW 05, Desa Sukawangi, Kec. Sukamakmur, Kab. Bogor', 'Cabai Keriting, Kacang Buncis, Tomat', 3, 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/Bogor+-+Mang+Eman.jpg'),
('Petani Utama', 'Mang Ali', '3201090507840003', 34, '085885863372',  'Menikah', 1,'Kampung Catang Malang 2 RT 03 RW 05, Desa Sukawangi, Kec. Sukamakmur, Kab. Bogor', 'Sawi Putih, Pakchoy, Daun Bawang', 10, 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/Bogor+-+Mang+Ali.jpg'),
('Petani Utama', 'Mang Peya', '3201090606820007', 36, NULL, 'Menikah', 2, 'Kampung Margaluyu, RT 01 RW 05, Desa Sukawangi, Kec. Sukamakmur, Kab. Bogor', 'Tomat, Sawi Putih', 8, 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/Bogor+-+Mang+Peyak.jpg'),
('Petani Utama', 'Dede Arifin', '3201090207000007', 19, '081546426521', 'Belum Menikah', 0, 'Kampung Cikuray, RT 02 RW 05, Desa Sukawangi, Kec. Sukamakmur, Kab. Bogor', 'Kacang Kapri', 5, 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/Bogor+-+Dede.jpg'),
('Petani Utama', 'Mang Majid', NULL, 55, NULL, 'Menikah', 0, 'Kampung Catang Malang 1 RT 04 RW 05, Desa Sukawangi, Kec. Sukamakmur, Kab. Bogor', 'Cabai Keriting, Tomat, Daun Bawang, Pakchoy, Kacang Kapri, Sawi Putih', 30, 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/Bogor+-+Mang+Majid.jpg'),
('Petani Utama', 'Mang Ujang', '3201091109830006', 35, NULL, 'Menikah', 2, 'Kampung Catang Malang 1 RT 04 RW 05, Desa Sukawangi, Kec. Sukamakmur, Kab. Bogor', 'Cabai Keriting, Tomat, Brokoli, Sawi Putih', 15, 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/Bogor+-+Mang+Ujang.jpg'),
('Petani Utama', 'Mang Uci', '3201091109830006', 39, NULL, 'Menikah', 3, 'Kampung Catang Malang 1 RT 04 RW 05, Desa Sukawangi, Kec. Sukamakmur, Kab. Bogor', 'Brokoli, Tomat, Sawi Putih, Pakcboy, Cabai Keriting', 25, 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/Bogor+-+Mang+Uci.jpg');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE farmers;