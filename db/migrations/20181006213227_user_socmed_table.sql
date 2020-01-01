
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE user_socmed(
    user_id integer REFERENCES users (id) ON DELETE restrict,
    facebook VARCHAR(100),
    instagram VARCHAR(100),
    linkedin VARCHAR(100),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE user_socmed;
