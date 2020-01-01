package model

import (
	"time"
)

type Land struct {
	ID          int
	Regional_id int
	Name        string
	Area        int
	Province    string
	City        string
	Address     string
	Postal_code string
	Ownership   string
	Certificate string
	survey_note string
	survey_file string
	Created_at  time.Time
	Update_at   time.Time
}
