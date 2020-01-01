
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TYPE investment_concurrent_status as ENUM ('Reviewed', 'Accepted', 'Rejected', 'Cancelled');

CREATE TABLE investment_concurrents (
    id SERIAL PRIMARY KEY,
    investor_id INTEGER NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
    project_id INTEGER NOT NULL REFERENCES projects (id) ON DELETE RESTRICT,
    concurrent_code VARCHAR (45),
    period INTEGER,
    quantity INTEGER,
    is_agree SMALLINT,
    status investment_concurrent_status,
    release_at TIMESTAMP NULL,
    expire_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE investment_concurrents;

DROP TYPE investment_concurrent_status;