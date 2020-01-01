
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE internships(
    id SERIAL PRIMARY KEY,
    name varchar (100),
    email VARCHAR (100),
    handphone VARCHAR (15),
    address TEXT,
    institute VARCHAR (100),
    major VARCHAR (50),
    gpa INTEGER,
    gpa_max INTEGER,
    specializations VARCHAR (191),
    first_position INTEGER REFERENCES job_positions (id) ON DELETE SET NULL,
    first_position_reason TEXT,
    second_position INTEGER REFERENCES job_positions (id) ON DELETE SET NULL,
    second_position_reason TEXT,
    days_of_intern INTEGER,
    date_start TIMESTAMP,
    reason TEXT,
    know_tanijoy VARCHAR (45),
    stage_id INTEGER REFERENCES career_stages (id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE internships;
