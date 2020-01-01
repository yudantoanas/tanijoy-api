package model

import (
	"time"
)

type Stage struct {
	ID          int
	Name        string
	Purpose     string
	Description string
	Created_at  time.Time
	Updated_at  time.Time
}
