package model

import (
	"time"
)

type Transaction struct {
	ID               int
	Transaction_code string
	Investor_id      int
	Related_id       int
	Amount           int
	Unique_code      int
	Type             string
	Note             string
	Payment_account  string
	Payment_method   string
	Transaction_at   time.Time
	Confirmed_at     time.Time
	Created_at       time.Time
	Updated_at       time.Time
}

type UnverifiedTransaction struct {
	Code         string
	Investor_id  int
	Project_id   int
	Total_amount int
	Bank_name    string
	Time         time.Time
}

type VerifiedTransaction struct {
	Message     string
	Concurrent  Investment_concurrent
	Transaction Transaction
}

type TransactionResult struct {
	Message     string
	Transaction Transaction
	Sub_total   int
	Unique_code int
	Total       int
}

type ConcurrentTransaction struct {
	Concurrent_id  int
	Investor_id    int
	Project_id     int
	Quantity       int
	Is_agree       int
	Status         string
	Amount         int
	Transaction_at time.Time
}

type TransactionHistory struct {
	Concurrent_transaction ConcurrentTransaction
	State                  string
}
