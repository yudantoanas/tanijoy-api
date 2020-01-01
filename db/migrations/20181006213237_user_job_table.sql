
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TYPE job_types AS enum ('Pegawai BUMN', 'Wiraswasta', 'Karyawan');
CREATE TYPE income_types AS enum ('A', 'B', 'C', 'Not Say');

-- A: < 5 Juta
-- B: 5 - 10 Juta
-- C: > 10 Juta

CREATE TABLE user_job(
    user_id integer REFERENCES users (id) ON DELETE restrict,
    npwp VARCHAR(30),
    job job_types,
    income income_types,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE user_job;

DROP TYPE job_types;
DROP TYPE income_types;