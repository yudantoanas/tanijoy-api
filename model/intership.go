package model

import (
	"time"
)

type Internship struct {
	ID                     int
	Name                   string
	Email                  string
	Handphone              string
	Address                string
	Institute              string
	Major                  string
	Gpa                    int
	Gpa_max                int
	Spelization            string
	First_position         int
	First_position_reason  string
	Second_position        int
	Second_position_reason string
	Day_of_intern          int
	Date_start             time.Time
	Reason                 string
	Know_tanijoy           string
	Stage_id               int
	Created_at             time.Time
	Updated_at             time.Time
}