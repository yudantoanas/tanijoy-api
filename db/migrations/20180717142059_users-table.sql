
-- +goose Up
CREATE TABLE users(
 id serial PRIMARY KEY,
 name VARCHAR (191) NOT NULL,
 email VARCHAR (191) UNIQUE,
 phone VARCHAR (20) NOT NULL,
 username VARCHAR (191) NOT NULL,
 password VARCHAR (191) NOT NULL,
 confirmed SMALLINT NOT NULL,
 confirmation_code VARCHAR (191) NOT NULL,
 api_token VARCHAR (191),
 remember_token VARCHAR (100),
 provider VARCHAR (191),
 provider_id VARCHAR (191),
 created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
 updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
 deleted_at TIMESTAMP
);

-- Seeder
INSERT INTO users (name, email, phone, username, password, confirmed, confirmation_code) VALUES 
('Febrian Admin', 'febrian@tanijoy.id', '082331012599', 'febrian', 'bismillah', 1, 'dummyuser'),
('Sofandi Arifin', 'sofandi@tanijoy.id', '08098999', 'sofandiarifin', 'landowner2017', 1, 'dummyuser'),
('Dwi Febaria', 'febaria@tanijoy.id', '08098999', 'dwifebaria', 'landowner2017', 1, 'dummyuser'),
('Ahmad Bakrie', 'bakrie@tanijoy.id', '08098999', 'ahmadbakrie', 'landowner2017', 1, 'dummyuser'),
('Alfian Riza', 'alfian@tanijoy.id', '080989999', 'alfianriza', 'alfiantanijoy2017', 1, 'dummyuser'),
('Ilham Akbar', 'ilham@tanijoy.id', '080989999', 'ilhamakbar', 'ilhamtanijoy2017', 1, 'dummyuser'),
('Kukuh Budi', 'kukuh@tanijoy.id', '085790916498', 'kukuh.budi', 'pmtanijoy2017', 1, 'dummyuser'),
('Jainal Rabin', 'jainal@tanijoy.id', '0812312311', 'jainal.rabin', 'rmmedantanijoy2017', 1, 'dummyuser');

-- +goose Down
DROP TABLE users;