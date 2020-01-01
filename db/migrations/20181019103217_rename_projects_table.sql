
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE IF EXISTS projects
RENAME TO proposals;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE IF EXISTS proposals
RENAME TO projects;
