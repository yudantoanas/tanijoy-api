
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE investments (
    id SERIAL PRIMARY KEY,
    investor_id INTEGER NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
    transaction_id INTEGER NOT NULL REFERENCES transactions (id) ON DELETE RESTRICT,
    land_id INTEGER REFERENCES lands (id) ON DELETE SET NULL,
    project_id INTEGER NOT NULL REFERENCES projects (id) ON DELETE RESTRICT,
    fieldmanager_id INTEGER REFERENCES users (id) ON DELETE SET NULL,
    farmer_id INTEGER REFERENCES farmers (id) ON DELETE SET NULL,
    stage_id INTEGER REFERENCES stages (id) ON DELETE SET NULL,
    phase_id INTEGER REFERENCES stages (id) ON DELETE SET NULL,
    investment_code VARCHAR (30),
    is_reinvest SMALLINT DEFAULT 0,
    expected_start_at TIMESTAMP,
    actual_start_at TIMESTAMP,
    expected_finish_at TIMESTAMP,
    actual_finish_at TIMESTAMP,
    delay INTEGER,
    sketch_path VARCHAR (191),
    certificate_path VARCHAR (191),
--     is_posted SMALLINT,
--     actual_return INTEGER,
--     actual_percentage INTEGER,
--     follow_up_at TIMESTAMP,
--     gift_at TIMESTAMP,
--     created_at TIMESTAMP DEFAULT current_timestamp,
--     updated_at TIMESTAMP DEFAULT current_timestamp,
--     deleted_at TIMESTAMP
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP,
    is_posted SMALLINT
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE investments;