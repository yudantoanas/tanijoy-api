package model

import "time"

type Commodity struct {
	ID         int
	Name       string
	Family     string
	Created_at time.Time
	Update_at  time.Time
}
