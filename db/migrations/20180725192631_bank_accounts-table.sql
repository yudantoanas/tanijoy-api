
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE bank_accounts (
    id SERIAL PRIMARY KEY,
    bank VARCHAR (50),
    alias VARCHAR (10),
    bank_code VARCHAR (5),
    account VARCHAR (20),
    holdername VARCHAR (50),
    branch VARCHAR (100),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

-- Seeder
INSERT INTO bank_accounts (bank, alias, bank_code, account, holdername, branch) VALUES 
('Bank Central Asia', 'BCA', '014', '4410070681', 'PT. Tanijoy Agriteknologi Nusantara', 'Jakarta Selatan'),
('Bank CIMB Niaga', 'CIMB', '022', '800155918300', 'PT. Tanijoy Agriteknologi Nusantara', 'Jakarta Selatan'),
('Bank CIMB Niaga Syariah', 'CIMB', '022', '860006960600', 'PT. Tanijoy Agriteknologi Nusantara', 'Jakarta Selatan'),
('Bank Negara Indonesia', 'BNI', '009', '0754920463', 'PT. Tanijoy Agriteknologi Nusantara', 'Jakarta Selatan'),
('Bank Mandiri', 'Mandiri', '008', '102 0007414 367', 'PT. Tanijoy Agriteknologi Nusantara', 'Jakarta Selatan'),
('Bank Rakyat Indonesia', 'BRI', '002', '0206.01.009043.30.2', 'PT. Tanijoy Agriteknologi Nusantara', 'BRI Kantor Cabang Khusus, Gd. BRI II');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE bank_accounts;
