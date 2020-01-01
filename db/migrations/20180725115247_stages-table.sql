
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE stages (
    id SERIAL PRIMARY KEY,
    name VARCHAR (100),
    purpose VARCHAR (100),
    description varchar (191),
    created_at TIMESTAMP DEFAULT current_timestamp,
    udpated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP 
);

-- Seeder
TRUNCATE TABLE stages;

INSERT INTO stages (name, purpose, description) values 
('pendanaan', 'Investment', 'Proyek dalam masa pendanaan.'),
('Menunggu', 'Investment', 'Proyek dalam masa menunggu hingga proyek sebelumnya selesai.'),
('Persiapan', 'Investment', 'Proyek dalam masa perancangan untuk menentukan timeline dan waktu mulai penanaman.'),
('Berjalan', 'Investment', 'Proyek budidaya sedang berjalan.'),
('Selesai', 'Investment', 'Proyek budidaya sudah berhasil dilaksanakan.'),
('Pembibitan', 'Planting', 'Fase pertumbuhan dari biji ke bibit.'),
('Vegetatif', 'Planting', 'Fase bertumbuhnya bibit dari yang ditandakan dengan bertambahnya tinggi dan jumlah daun.'),
('Generatif', 'Planting', 'Fase perkembangan yang ditandai dengan munculnya bunga dan selanjutnya buah.'),
('Panen', 'Planting', 'Fase perkembangan yang ditandai dengan munculnya bunga dan selanjutnya buah.'),
('Bagi Hasil', 'Planting', 'Perhitungan dan pembagian dari hasil panen.'),
('Pengolahan Lahan', 'Planting', 'Pengolahan lahan untuk menjadi lahan yang siap ditanam.');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE stages;
