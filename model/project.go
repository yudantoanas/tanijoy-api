package model

import (
	"time"
)

type Project struct {
	ID              int
	Commodities_id  int
	Name            string
	Amount          int
	Return_expected int
	Return_min      int
	Return_max      int
	Area            int
	Period          int
	Planting_system string
	Plants_count    int
	Risk            string
	Image_path      string
	Thumbnail_path  string
	Is_featured     int
	Created_at      time.Time
	Update_at       time.Time
}

type Project_descriptions struct {
	ID         int
	Project_id int
	Label      string
	Content    string
	Order      int
	Created_at time.Time
	Update_at  time.Time
}

type Project_files struct {
	ID          int
	Project_id  int
	Uploaded_by int
	Filename    string
	Path        string
	Label       string
	Purpose     string
	Created_at  time.Time
	Update_at   time.Time
}

type Project_risks struct {
	ID         int
	Project_id int
	Risk       string
	Mitigation string
	Created_at time.Time
	Update_at  time.Time
}

type Land_project struct {
	ID         int
	Project_id int
	Land_id    int
	Quantity   int
	Hold       int
	Ready      int
}

func (Land_project) TableName() string {
	return "land_project"
}

func (Project) TableName() string {
	return "proposals"
}

type ProjectStatus struct {
	Project_id int
	Jumlah     int
}
