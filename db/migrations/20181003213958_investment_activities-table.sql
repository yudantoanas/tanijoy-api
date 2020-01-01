
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE investment_activities(
  id SERIAL PRIMARY KEY,
  investment_id INTEGER NOT NULL REFERENCES investments (id) ON DELETE RESTRICT,
  activity_category_id INTEGER NOT NULL REFERENCES activity_categories (id) ON DELETE RESTRICT,
  activity_by INTEGER NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
  activity VARCHAR(45) NOT NULL,
  related_id INTEGER NOT NULL,
  related_to VARCHAR(45) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  deleted_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE investment_activities;

