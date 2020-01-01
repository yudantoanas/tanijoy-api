package model

import (
	"time"
)

type Regional struct {
	ID         int
	Manager_id int
	Name       string
	Created_at time.Time
	Update_at  time.Time
}
