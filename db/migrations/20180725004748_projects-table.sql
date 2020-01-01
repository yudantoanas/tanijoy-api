
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TYPE project_risk_parameters as ENUM ('Rendah', 'Sedang', 'Tinggi');

CREATE TYPE project_types as ENUM ('Budidaya', 'Trading');

CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    commodities_id INTEGER REFERENCES commodities (id) ON DELETE SET NULL,
    name VARCHAR(100),
    type project_types,
    amount INTEGER,
    return_expected INTEGER,
    return_min INTEGER,
    return_max INTEGER,
    area INTEGER,
    period INTEGER,
    plants_count INTEGER,
    risk project_risk_parameters,
    image_path VARCHAR (191),
    thumbnail_path VARCHAR (191),
    is_featured SMALLINT,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

-- seeder
TRUNCATE TABLE projects;

INSERT INTO projects (commodities_id, name, type, amount, return_expected, return_min, return_max, area, period, plants_count, risk, image_path, thumbnail_path, is_featured) VALUES
(1, 'Cabe Keriting Bogor', 'Budidaya', 22275000, 18, 5, 25, 2500, 180, 6000, 'Tinggi', '', '',0),
(3, 'Brokoli Bogor', 'Budidaya', 2375000, 10, 3, 12, 2000, 120, 5500, 'Sedang', '', '',1),
(2, 'Daun Bawang Bogor', 'Budidaya', 1945000, 6, 3, 12, 1000, 90, 7000, 'Rendah', '', '',0),
(4, 'Kacang Kapri Bogor', 'Budidaya', 11375000, 10, 3, 12, 2000, 120, 2000, 'Sedang', '', '',0),
(5, 'Pak Choy Bogor', 'Budidaya', 3400000, 6, 3, 9, 1000, 90, 4000, 'Rendah', '', '',0),
(6, 'Sawi Putih Bogor', 'Budidaya', 3092000, 6, 3, 9, 1000, 90, 4000, 'Rendah', '', '',0);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE projects;
DROP TYPE project_risk_parameters;
DROP TYPE project_types;