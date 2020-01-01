package model

import (
	"time"
)

type Investment struct {
	ID                 int
	Investor_id        int
	Transaction_id     int
	Project_id         int
	Land_id            int
	Fieldmanager_id    int
	Farmer_id          int
	Stage_id           int
	Phase_id           int
	Investment_code    string
	Is_reinvest        int
	Expected_start_at  time.Time
	Expected_finish_at time.Time
	Delay              int
	Sketch_path        string
	Certificate_path   string
	Created_at         time.Time
	Updated_at         time.Time
}

type Investment_concurrent struct {
	ID              int
	Investor_id     int
	Project_id      int
	Concurrent_code string
	Period          int
	Quantity        int
	Is_agree        int
	Status          string
	Created_at      time.Time
	Updated_at      time.Time
	Expire_at       time.Time
}

type Investment_review struct {
	ID                 int
	Investment_id      int
	Satisfaction_level string
	Review             string
	Is_shared          int
	Created_at         time.Time
}

type Investment_update struct {
	ID                    int
	Project_id            int
	Related_id            int
	Category              string
	Aditional_information string
	Created_at            time.Time
	Updated_at            time.Time
}

type Investment_closing struct {
	ID                int
	Investment_id     int
	Statement         string
	File_path         string
	Actual_return     int
	Actual_percentage int
	Created_at        time.Time
}

type Investment_activity struct {
	ID                   int
	Investment_id        int
	Activity_category_id int
	Activity_by          int
	Activity             string
	Related_id           int
	Related_to           string
	Created_at           time.Time
	Updated_at           time.Time
}

type Investment_gallery struct {
	ID                     int
	Investment_id          int
	Project_id             int
	Investment_activity_id int
	Src                    string
	Created_at             time.Time
	Updated_at             time.Time
}

type Investment_file struct {
	ID                     int
	Investment_id          int
	Investment_activity_id int
	Title                  string
	Path                   string
	Created_at             time.Time
	Updated_at             time.Time
}

type ListInvestmentPhoto struct {
	Photos      []string
	More_photos int
}

type ActivitiesNotification struct {
	Latest_update int
	Invest_update int
}

type UserConcurrent struct {
	Message     string
	Concurrent  Investment_concurrent
	Sub_total   int
	Unique_code int
	Total       int
}

func (Investment_activity) TableName() string {
	return "investment_activities"
}

func (Investment_gallery) TableName() string {
	return "investment_galleries"
}

func (Investment_update) TableName() string {
	return "investment_update"
}

func (Investment_closing) TableName() string {
	return "investment_closing"
}

func (Investment_review) TableName() string {
	return "investment_review"
}
