package models

import "time"

type Costumer struct {
	ID             uint   `gorm:"primary_key"`
	Name           string `gorm:"type:varchar(100);unique_index"`
	Description    string `gorm:"type:text"`
	Representative string
	Location       string
	Email          string `gorm:"type:varchar(100);unique_index"`
	Tel            string
	Fax            string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
