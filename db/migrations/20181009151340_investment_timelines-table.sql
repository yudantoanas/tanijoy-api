
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE investment_timelines(
  id SERIAL PRIMARY KEY,
  investment_id INTEGER NOT NULL REFERENCES investments (id) ON DELETE RESTRICT,
  pembibitan_start TIMESTAMP DEFAULT current_timestamp,
  pembibitan_end TIMESTAMP DEFAULT current_timestamp,
  vegetatif_start TIMESTAMP DEFAULT current_timestamp,
  vegetatif_end TIMESTAMP DEFAULT current_timestamp,
  generatif_start TIMESTAMP DEFAULT current_timestamp,
  generatif_end TIMESTAMP DEFAULT current_timestamp,
  panen_start TIMESTAMP DEFAULT current_timestamp,
  panen_end TIMESTAMP DEFAULT current_timestamp,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE investment_timelines;

