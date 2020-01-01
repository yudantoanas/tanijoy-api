
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO investment_concurrents (investor_id, project_id, concurrent_code, period, quantity, is_agree, status, expire_at, release_at) VALUES
(9, 1, 'lorem001', 180, 1, 1, 'Accepted', null, null),
(9, 1, 'lorem0002', 180, 2, 1, 'Cancelled', '02/11/2018 10:55:00', '02/11/2018 11:05:00'),
(9, 1, 'lorem0003', 180, 1, 1, 'Accepted', null, null);

INSERT INTO land_concurrents (concurrent_id, land_id, project_id, quantity) VALUES
(1, 1, 1, 1),
(2, 1, 1, 2),
(3, 1, 1, 1);

INSERT INTO transactions (transaction_code, investor_id, related_id, amount, unique_code, type, note, payment_account, payment_method, transaction_at, confirmed_at) VALUES
('trf-lorem0001', 9, 1, 2275000, 456, 'Send', '', 'BCA', 'Transfer', '2018-05-03 10:55:00', '2018-05-03 11:55:00'),
('trf-lorem0002', 9, 2, 4550000, 123, 'Send', '', 'BCA', 'Transfer', null , null),
('trf-lorem0003', 9, 3, 2275000, 234, 'Send', '', 'BCA', 'Transfer', '2018-11-03 10:55:00' , '2018-11-03 11:55:00');

INSERT INTO investments (investor_id, transaction_id, land_id, project_id, fieldmanager_id, farmer_id, stage_id, phase_id, investment_code, is_reinvest, expected_start_at, actual_start_at, expected_finish_at, actual_finish_at) VALUES
(9, 1, 1, 1, 5, 1, 5, 10, 'CKB1805001', 1, '2018-05-10', '2018-05-10', '2018-10-10', '2018-10-17'),
(9, 2, 1, 1, 5, 1, 4, 11, 'CKB1811001', 1, '2018-11-09', '2018-11-09', '2019-05-08', null );

ALTER TABLE harvests
ADD COLUMN harvest_date TIMESTAMP,
ADD COLUMN type VARCHAR (45),
ADD COLUMN total INT;

INSERT INTO harvests (investment_id, harvest_date, type, qty, price, total) VALUES
(1, '2018-09-19', 'Utama', 100, 8000, 800000),
(1, '2018-09-19', 'Tumpang Sari', 25, 2000, 50000),
(1, '2018-09-26', 'Utama', 50, 10000, 500000),
(1, '2018-10-03', 'Utama', 70, 12000, 840000);

TRUNCATE TABLE activity_categories RESTART IDENTITY cascade;

INSERT INTO activity_categories (category) VALUES
('Pembukaan & Pengolahan Lahan'),
('Penanaman'),
('Pemupukan'),
('Penyemprotan Nutrisi & Pestisida'),
('Penyiangan'),
('Pemangkasan'),
('Pemanenan'),
('Lainnya');

ALTER TABLE investment_activities ALTER COLUMN related_id DROP NOT NULL;
ALTER TABLE investment_activities ALTER COLUMN related_to DROP NOT NULL;

INSERT INTO investment_activities (investment_id, activity_category_id, activity_by, activity) VALUES
(1, 1, 5, 'Pembukaan Lahan'),
(1, 2, 5, 'Penanaman Bibit Cabe Keriting'),
(1, 8, 5, 'Penanaman Ajir'),
(1, 3, 5, 'Pemupukan Pupuk Kandang'),
(1, 4, 5, 'Penyemprotan Nutrisi & Pestisida'),
(1, 3, 5, 'Pemupukan Pupuk Kimia'),
(1, 5, 5, 'Pencabutan rumput liar'),
(1, 6, 5, 'Pemangkasan daun'),
(1, 7, 5, 'Panen 100 Kg Cabe Keriting'),
(1, 7, 5, 'Panen 25 Kg Daun Bawang'),
(1, 7, 5, 'Panen 50 Kg Cabe Keriting'),
(1, 7, 5, 'Panen 50 Kg Cabe Keriting');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back