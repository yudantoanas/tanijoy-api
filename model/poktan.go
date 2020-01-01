package model

import (
	"time"
)

type Poktan struct {
	ID            int
	Name          string
	Email         string
	Phone         string
	Address       string
	Area          int
	Farmers_count int
	Vegetables    string
	Stage_id      int
	Created_at    time.Time
	Updated_at    time.Time
}

type UnverifiedPoktan struct {
	Name         string
	Email        string
	Phone        string
	Address      string
	Vegetables   string
	Area_width   int
	Total_farmer int
	Register_at  time.Time
}

func (Poktan) TableName() string {
	return "poktan_candidates"
}
