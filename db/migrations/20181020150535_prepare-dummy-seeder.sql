
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO farmer_commodity (farmer_id, commodity_id) VALUES
(1,2),
(2,2);

INSERT INTO proposal_stage (stage_id, proposal_id, initial_time_span, final_time_span, days_of_delayed) VALUES
(1, 3, 35, 45, 0);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

