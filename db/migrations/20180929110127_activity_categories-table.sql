
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE activity_categories(
  id serial PRIMARY KEY,
  category VARCHAR(45) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  deleted_at TIMESTAMP
);

INSERT INTO activity_categories (category) VALUES
('Pemupukan'), ('Pembukaan lahan'), ('Penyiangan'), ('Panen');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE activity_categories;
