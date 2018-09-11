package models

import "time"

type Immigration struct {
	ID        uint `gorm:"primary_key" form:"id" json:"id"`
	StartDate time.Time
	EndDate   time.Time
	Visa      bool
	Status    string `sql:"type:enum('inProgress','completed','canceled');DEFAULT:'inProgress'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
