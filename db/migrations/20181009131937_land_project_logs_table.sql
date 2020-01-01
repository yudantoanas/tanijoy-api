
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TYPE plus_minus_operators AS enum ('plus', 'minus');
CREATE TYPE land_project_components AS enum ('quantity', 'ready');

-- Add primary key to land_project if doesn't exist
-- ALTER TABLE land_project
-- 	ADD COLUMN IF NOT EXISTS id SERIAL PRIMARY KEY;

CREATE TABLE land_project_logs(
    id SERIAL PRIMARY KEY,
    land_project_id integer REFERENCES land_project (id) ON DELETE SET NULL,
    operator plus_minus_operators,
    component land_project_components,
    qty INTEGER,
    reason TEXT,
    created_at TIMESTAMP DEFAULT current_timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

-- drop primary key
ALTER TABLE land_project
	DROP CONSTRAINT land_project_pkey,
	DROP COLUMN id;

-- drop table
DROP TABLE land_project_logs;

-- drop type
DROP TYPE plus_minus_operators;
DROP TYPE land_project_components;