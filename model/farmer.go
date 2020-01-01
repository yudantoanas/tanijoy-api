package model

import (
	"time"
)

type Farmer struct {
	ID                   int
	Name                 string
	Phone                string
	Address              string
	Age                  int
	Year_of_experiences int
	Is_married           int
	Number_of_childs     int
	Specializations      string
	Profpic_src          string
	Account_id           int
	Created_at           time.Time
	Update_at            time.Time
}
