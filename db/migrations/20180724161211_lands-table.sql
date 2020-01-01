
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TYPE land_ownerships as ENUM ('Pribadi', 'Milik Bersama');
CREATE TYPE land_certificates as ENUM ('Bersertifikat', 'Tidak Bersertifikat');

CREATE TABLE lands (
    id SERIAL PRIMARY KEY,
    regional_id INTEGER REFERENCES regionals (id) ON DELETE RESTRICT,
    landowner_id INTEGER REFERENCES users (id) ON DELETE RESTRICT,
    name VARCHAR (100),
    area INTEGER,
    province VARCHAR (50),
    city VARCHAR (50),
    address TEXT,
    postal_code VARCHAR (10),
    ownership land_ownerships,
    certificate land_certificates,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

-- Seeder
TRUNCATE TABLE lands;

INSERT INTO lands (regional_id, landowner_id, name, area, province, city, address, postal_code, ownership, certificate) VALUES 
(1, 2, 'Cetang Malang', 100000, 'Jawa Barat', 'Kabupaten Bogor', 'Kp. Cetang Malang, ds Sukawangi, kec. Sukamakmur, Kab. Bogor', 123123, 'Pribadi', 'Bersertifikat'),
(1, 3, 'Cikuray', 40000, 'Jawa Barat', 'Kabupaten Bogor', 'Kp. Cikuray, Ds. Sukawangi, Kec. Sukamakmur, Kab. Bogor', 123123, 'Pribadi', 'Bersertifikat'),
(1, 4, 'Neglasari', 100000, 'Jawa Barat', 'Kabupaten Bogor', 'Kp. Neglasari, Ds. Sukawangi, Kec. Sukamakmur, Kab. Bogor', 123123, 'Pribadi', 'Bersertifikat');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE lands;

DROP TYPE land_ownerships;

DROP TYPE land_certificates;
