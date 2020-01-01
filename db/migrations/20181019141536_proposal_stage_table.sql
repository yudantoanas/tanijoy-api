
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE proposal_stage (
    id SERIAL PRIMARY KEY,
    stage_id INTEGER REFERENCES stages (id) ON DELETE SET NULL,
    proposal_id INTEGER REFERENCES proposals (id) ON DELETE SET NULL,
    initial_time_span INTEGER,
    final_time_span INTEGER,
    days_of_delayed INTEGER,
    created_at TIMESTAMP default current_timestamp ,
    udpated_at TIMESTAMP default current_timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE proposal_stage;