
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
TRUNCATE TABLE farmers RESTART IDENTITY CASCADE;

INSERT INTO farmers (farmer_type, name, year_of_experiences, specializations, profpic_src) VALUES
('Petani Utama', 'Mang Aep', 18, 'Brokoli,Tomat,Kacang Kapri,Pak Choy','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/brokoli/Mang+Aep_+Brokoli%2C+Kacang+Kapri%2C+Pak+Choy_+18+tahun.jpg'),
('Petani Utama', 'Mang Deni', 10, 'Brokoli,Cabai Keriting,Tomat,Daun Bawang','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/brokoli/Mang+Deni_+Brokoli%2C+Tomat%2C+Daun+Bawang_+10+Tahun.jpg'),
('Petani Utama', 'Mang Oman', 10, 'Brokoli,Cabai Keriting,Tomat','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/brokoli/Mang+Oman_+Brokoli%2C+Cabai+Keriting%2C+Tomat_10+Tahun.jpg'),
('Petani Utama', 'Mang Peyak', 8, 'Brokoli,Cabai Keriting,Tomat,Sawi Putih','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/brokoli/Mang+Peyak_+Brokoli%2C+Cabai+Keriting%2C+Tomat%2C+Sawi_+8+Tahun.jpg'),
('Petani Utama', 'Mang Ujang', 15, 'Brokoli,Cabai Keriting,Tomat,Sawi Putih','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/brokoli/Mang+Ujang_+Brokoli%2C+Cabai+Keriting%2C+Tomat%2C+Sawi+Putih_+15+Tahun+.jpg'),
('Petani Utama', 'Abah Sarmin', 30, 'Cabai Keriting,Tomat,Sawi Putih, Pak Choy','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/cabai-keriting/Abah+Sarmin_+Cabai+Keriting%2CSawi+Putih%2C+Tomat%2C+Pak+Cho_+30+tahun+.jpg'),
('Petani Utama', 'Mang Ali', 10, 'Cabai Keriting,Tomat,Sawi Putih,Daun Bawang,Pak Choy','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/cabai-keriting/Mang+Ali_+Cabai+Keriting%2C+Sawi+Putih%2C+Pak+Choy%2C+Daun+Bawang%2C+Tomat_+10+Tahun.jpg'),
('Petani Utama', 'Mang Maman', 36, 'Cabai Keriting','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/cabai-keriting/Mang+Maman_+Cabai+Keriting_36+Tahun.jpg'),
('Petani Utama', 'Mang Udan', 4, 'Cabai Keriting,Daun Bawang','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/cabai-keriting/Mang+Udan_+Cabai+Keriting%2C+Daun+Bawang_+4+tahun.jpg'),
('Petani Utama', 'Mang Ujan', 15, 'Brokoli,Cabai Keriting,Tomat,Sawi Putih','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/cabai-keriting/Mang+Ujan_+Cabai+Keriting%2C+Tomat%2C+Brokoli%2C+Sawi+Putih+_+15+Tahun.jpg'),
('Petani Utama', 'Mang Yadi', 10, 'Cabai Keriting,Tomat','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/cabai-keriting/Mang+Yadi_+Cabai+Keriting%2C+Tomat_+10+Tahun+.jpg'),
('Petani Utama', 'Mang Eman', 5, 'Daun Bawang,Cabai Keriting,Tomat,Kacang Buncis','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/daun-bawang/Mang+Eman_+Daun+Bawang%2C+Cabai+Keriting%2C+Kacang+Buncis%2CTomat_+5+Tahun.jpg'),
('Petani Utama', 'Mang Majid', 30, 'Daun Bawang,Cabai Keriting,Tomat,Pak Choy,Kacang Kapri,Sawi Putih','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/daun-bawang/Mang+Majid_+Daun+Bawang%2C+Cabai+Keriting%2C+Tomat%2C+Pak+Choy%2C+Kacang+Kapri%2C+Sawi+Putih_30+Tahun.jpg'),
('Petani Utama', 'Mang Soleh', 30, 'Daun Bawang,Cabai Keriting,Pak Choy,Sawi Putih','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/daun-bawang/Mang+Soleh_+Daun+Bawang%2C+Cabai+Keriting%2C+Pak+Choy%2C+Sawi+Putih_+30+Tahun.jpg'),
('Petani Utama', 'Dede', 5, 'Kacang Kapri,Tomat','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/kacang-kapri/Dede_+Kacang+Kapri%2CTomat_+5+Tahun.jpg'),
('Petani Utama', 'Mang Anang', 8, 'Kacang Kapri,Cabai Keriting,Pak Choy,Kacang Buncis,Daun Bawang','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/kacang-kapri/Mang+Anang_Kacang+Kapri%2C+Cabai+Keriting%2C+Pak+Choy%2C+Kacang+Buncis%2C+Daun+Bawang_+8+Tahun.jpg'),
('Petani Utama', 'Mang Asep', 20, 'Kacang Kapri,Pak Choy,Sawi Putih,Terong,Tomat','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/kacang-kapri/Mang+Asep_Kacang+Kapri%2C+Pak+Choy%2C+Sawi+Putih%2C+Terong%2C+Tomat_20+Tahun.jpg'),
('Petani Utama', 'Mang Oleh', 25, 'Kacang Kapri,Daun Bawang,Buncis,Cabai Keriting','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/kacang-kapri/Mang+Oleh_Kacang+Kapri%2C+Daun+Bawang%2C+Buncis%2C+Cabai+Keriting_+25+Tahun.jpg'),
('Petani Utama', 'Bu Fitri', 20, 'Kentang Granola','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/kentang-granola/Bu+Fitri_+Kentang+Granola_20+Tahun.jpeg'),
('Petani Utama', 'Pak Deswan', 10, 'Kentang Granola','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/kentang-granola/Pak+Deswan-+Kentang+Granola-+10+Tahun+.jpeg'),
('Petani Utama', 'Pak Munthe', 15, 'Kentang Granola','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/kentang-granola/Pak+Munthe_+Kentang+Granola_+15+Tahun+.jpeg'),
('Petani Utama', 'Pak Saragih', 20, 'Kentang Granola','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/kentang-granola/Pak+Saragih_+Kentang+Granola_+20+Tahun+.jpeg'),
('Petani Utama', 'Mang Soleh', 20, 'Pak Choy,Cabai Keriting,Daun Bawang,Sawi Putih','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/pak-choy/Mang+Soleh_+Pak+Choy%2C+Cabai+Keriting%2C+Daun+Bawang%2C+Sawi+Putih_+20+Tahun.jpg'),
('Petani Utama', 'Mang Uci', 20, 'Pak Choy,Brokoli,Tomat,Sawi Putih,Cabai Keriting','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/pak-choy/Mang+Uci_+Pak+Choy%2C+Brokoli%2C+Tomat%2C+Sawi+Putih%2C+Cabai+Keriting%2C+20+Tahun+.jpg'),
('Petani Utama', 'Mang Majid', 30, 'Sawi Putih,Cabai Keriting,Tomat,Daun Bawang,Pak Choy,Kacang Kapri','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/sawi-putih/Mang+Majid_+Sawi+Putih%2C+Cabai+Keriting%2C+Tomat%2C+Daun+Bawang%2C+Pak+Choy%2C+Kacang+Kapri_+30+Tahun.jpg'),
('Petani Utama', 'Mang Ujang', 15, 'Sawi Putih,Cabai Keriting,Tomat,Brokoli','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/sawi-putih/+Mang+Ujang_+Sawi+Putih%2C+Cabai+keriting%2C+Tomat%2C+Brokoli_+15+Tahun.jpg'),
('Petani Utama', 'Mang Asep', 15, 'Tomat,Kacang Kapri,Pak Choy,Sawi Putih,Terong','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/tomat-bogor/Mang+Asep_+Tomat%2CKacang+Kapri%2C+Pak+Choy%2C+Sawi+Putih%2C+Terong_20+Tahun.jpg'),
('Petani Utama', 'Mang Ba`ah', 15, 'Tomat','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/tomat-bogor/Mang+Ba%C2%B4ah_Tomat_.png'),
('Petani Utama', 'Mang Pian', 15, 'Tomat,Brokoli','https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/farmers/tomat-bogor/Mang+Pian_+Tomat.png');

INSERT INTO farmer_commodity (farmer_id, commodity_id) VALUES
(1,3),(1,4),(1,5),(1,9),
(2,3),(2,1),(2,9),(2,2),
(3,3),(3,1),(3,9),
(4,3),(4,1),(4,9),(4,6),
(5,3),(5,1),(5,9),(5,6),
(6,1),(6,9),(6,6),(6,5),
(7,1),(7,9),(7,6),(7,2),(7,5),
(8,1),
(9,1),(9,2),
(10,3),(10,1),(10,9),(10,6),
(11,1),(11,9),
(12,2),(12,1),(12,9),(12,4),
(13,2),(13,1),(13,9),(13,5),(13,6),
(14,2),(14,1),(14,5),(14,6),
(15,4),(15,9),
(16,4),(16,1),(16,5),(16,2),
(17,4),(17,5),(17,6),(17,9),
(18,4),(18,2),(18,1),
(19,7),(20,7),(21,7),(22,7),
(23,5),(23,1),(23,2),(23,6),
(24,5),(24,3),(24,9),(24,6),(24,1),
(25,6),(25,1),(25,9),(25,2),(25,5),(25,4),
(26,6),(26,1),(26,9),(26,3),
(27,9),(27,4),(27,5),(27,6),
(28,9),(29,9),(29,3);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

