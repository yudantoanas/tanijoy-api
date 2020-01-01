package model

import (
	"time"
)

// user model
type User struct {
	ID                int
	Name              string
	Username          string
	Password          string
	Email             string
	Phone             string
	Provider          string
	Provider_id       string
	Confirmed         int
	Confirmation_code string
	Api_token         string
	Remember_token    string
	Created_at        time.Time
	Updated_at        time.Time
}

type User_information struct {
	User_id           int
	Identities_number string
	Province          string
	City              string
	Address           string
	Postal_code       int
	Gender            string
	Profile_progress  int
	Birth_at          time.Time
	Created_at        time.Time
	Updated_at        time.Time
}

type User_file struct {
	ID         int
	User_id    int
	Purpose    string
	Type       string
	Path       string
	Created_at time.Time
	Updated_at time.Time
}

type User_bank struct {
	ID             int
	User_id        int
	Bank           string
	Branch         string
	Account_number string
	Holdername     string
	Is_primary     int
	Created_at     time.Time
	Updated_at     time.Time
}

type User_job struct {
	User_id    int
	Npwp       string
	Job_type   string
	Grade      string
	Income     string
	Created_at time.Time
	Updated_at time.Time
}

type User_socmed struct {
	User_id   int
	Facebook  string
	Instagram string
	Linkedin  string
}

type Role_user struct {
	User_id    int
	Role_id    int
	Created_at time.Time
	Updated_at time.Time
}

type UserAuth struct {
	Message string
	User    User
}

func (Role_user) TableName() string {
	return "role_user"
}

func (User_socmed) TableName() string {
	return "user_socmed"
}

func (User_job) TableName() string {
	return "user_job"
}

func (User_information) TableName() string {
	return "user_information"
}
