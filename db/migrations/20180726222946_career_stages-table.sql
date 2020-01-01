
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE career_stages (
    id SERIAL PRIMARY KEY,
    name VARCHAR (50),
    description varchar (191),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

-- Seeder
INSERT INTO career_stages (name, description) VALUES 
('Applied', 'Applicants memasukkan data dan belum di proses'), 
('Culture Interview', 'Interview tahap pertama soal kecocokan dengan culture company'), 
('Technical Interview', 'Interview tahap kedua soal kemampuan'),
('Review', 'Proses review dari founder'),
('Accepted', 'Applicants diterima'),
('Rejected', 'Applicants ditolak');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE career_stages;
