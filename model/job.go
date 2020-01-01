package model

import (
	"time"
)

type Job_position struct {
	ID           int
	Name         string
	Description  string
	Requirements string
	Label        string
	Created_at   time.Time
	Updated_at   time.Time
	Category_id  int
}