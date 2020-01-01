
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE investment_galleries(
  id SERIAL PRIMARY KEY,
  investment_id INTEGER NOT NULL REFERENCES investments (id) ON DELETE RESTRICT,
  project_id INTEGER NOT NULL REFERENCES projects (id) ON DELETE RESTRICT,
  investment_activity_id INTEGER NOT NULL REFERENCES investment_activities (id) ON DELETE RESTRICT,
  src VARCHAR (100),
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE investment_galleries;

