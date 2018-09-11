package models

import (
	"time"
)

type Relocation struct {
	ID                    uint `gorm:"primary_key" form:"id" json:"id"`
	StartDate             time.Time
	EndDate               time.Time
	Accomodated           bool
	Address               string
	TelephoneSubscriber   bool
	ElectricitySubscriber bool
	Status                string `sql:"type:enum('waiting','inProgress','completed','canceled');DEFAULT:'waiting'"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
