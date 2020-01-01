
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE investment_files(
  id SERIAL PRIMARY KEY,
  investment_id INTEGER NOT NULL REFERENCES investments (id) ON DELETE RESTRICT,
  investment_activity_id INTEGER NOT NULL REFERENCES investment_activities (id) ON DELETE RESTRICT,
  title VARCHAR (191) NOT NULL,
  path VARCHAR(191) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE investment_files;

