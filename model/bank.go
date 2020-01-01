package model

import (
	"time"
)

type Bank_account struct {
	ID         int
	Bank       string
	Alias      string
	Bank_code  string
	Account    string
	Holdername string
	Branch     string
	Created_at time.Time
	Updated_at time.Time
}

type Bank_list struct {
	ID         int
	Name       string
	Alias      string
	Logo_path  string
	Created_at time.Time
	Updated_at time.Time
}

type Payment_bank struct {
	Account   string
	Logo_path string
}

func (Bank_list) TableName() string {
	return "bank_list"
}
