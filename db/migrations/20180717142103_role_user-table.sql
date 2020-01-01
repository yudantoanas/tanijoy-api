
-- +goose Up
CREATE TABLE role_user(
 user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE restrict,
 role_id INTEGER NOT NULL REFERENCES roles (id) ON DELETE restrict,
 created_at TIMESTAMP DEFAULT current_timestamp,
 updated_at TIMESTAMP DEFAULT current_timestamp,
 deleted_at TIMESTAMP
);

-- seeder
INSERT INTO role_user (user_id, role_id) VALUES
(1, 1), (2, 2), (3, 2), (4, 2),
(5, 7), (6, 7), (7, 6), (8, 6);

-- +goose Down
DROP TABLE role_user;