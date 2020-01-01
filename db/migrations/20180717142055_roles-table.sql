
-- +goose Up
CREATE TABLE roles(
 id serial PRIMARY KEY,
 name VARCHAR (191) NOT NULL,
 description VARCHAR (191) NOT NULL,
 created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
 updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
 deleted_at TIMESTAMP
);

-- Seeder 
INSERT INTO roles (name, description) VALUES 
('Admin', 'A guy who can give you permission. Be kind with them.'),
('Investor', 'Someone who give their capital to invest on our project. They have a the rights to get a periodically reports from their invested projects.'),
('Landowner', 'Someone who give their capital to invest on our project. They have a the rights to get a periodically reports from their invested projects.'),
('Customer Service', 'Front line who answer and response all the investor question and ask.'),
('Project Manager', 'Responsible for all projects going well, watch your step bro.'),
('Regional Manager', 'Responsible with field manager and projects in specific areas.'),
('Field Manager', 'Someone who responsible with ongoing project, create pre-planting project and can cover up to 5 hectars. Also create daily activities and weekly activities for periodically report.'),
('Farmer', 'Hero who make foods for us.');

-- +goose Down
DROP TABLE roles;