
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TYPE project_labels as ENUM ('image', 'document');

CREATE TYPE project_purposes as ENUM ('prospektus', 'gallery');

CREATE TABLE project_files (
    id SERIAL PRIMARY KEY,
    project_id INTEGER NOT NULL REFERENCES projects (id) ON DELETE CASCADE,
    uploaded_by INTEGER REFERENCES users (id) ON DELETE RESTRICT,
    filename VARCHAR (100),
    path VARCHAR (191),
    label project_labels,
    purpose project_purposes,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

-- Seeder
TRUNCATE TABLE project_files;

INSERT INTO project_files (project_id, filename, path, label, purpose) VALUES
(1, 'Prospektus Cabe Keriting Bogor', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/documents/Dokumen-Investasi-Cabai-Keriting-Bogor.pdf', 'document', 'prospektus'),
(2, 'Prospektus Brokoli Bogor', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/documents/Dokumen-Investasi-Brokoli-Bogor.pdf', 'document', 'prospektus'),
(3, 'Prospektus Duan Bawang Bogor', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/documents/Dokumen-Investasi-Daun-Bawang-Bogor.pdf', 'document', 'prospektus'),
(4, 'Prospektus Kacang Kapri Bogor', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/documents/Dokumen-Investasi-Kacang-Kapri-Bogor.pdf', 'document', 'prospektus'),
(5, 'Prospektus Pak Choy Bogor', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/documents/Dokumen-Investasi-Pak-Choy-Bogor.pdf', 'document', 'prospektus'),
(6, 'Prospektus Sawi Putih Bogor', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/documents/Dokumen-Investasi-Sawi-Putih-Bogor.pdf', 'document', 'prospektus'),
(1, 'thumbnail-cabe-keriting-rectangle', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-cabe-keriting-rectangle.jpg', 'image', 'gallery'),
(2, 'thumbnail-brokoli-rectangle', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-brokoli-rectangle.jpg', 'image', 'gallery'),
(3, 'thumbnail-daun-bawang-rectangle', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-daun-bawang-rectangle.jpg', 'image', 'gallery'),
(4, 'thumbnail-kacang-kapri-rectangle', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-kacang-kapri-rectangle.jpg', 'image', 'gallery'),
(5, 'thumbnail-pak-choy-rectangle', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-pak-choy-rectangle.jpg', 'image', 'gallery'),
(6, 'thumbnail-sawi-putih-rectangle', 'https://s3-ap-southeast-1.amazonaws.com/tanijoys3sg/projects/images/thumbnail-sawi-putih-rectangle.jpg', 'image', 'gallery');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE project_files;

DROP TYPE project_purposes;
DROP TYPE project_labels;
