package model

import (
	"time"
)

type Activity_category struct {
	ID         int
	Category   string
	Created_at time.Time
	Updated_at time.Time
}
