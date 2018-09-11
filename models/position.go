package models

import "time"

type Position struct {
	ID           uint   `gorm:"primary_key" form:"id" json:"id"`
	Title        string `gorm:"type:varchar(100);index"`
	Description  string `gorm:"type:text"`
	Location     string `gorm:"type:varchar(255)"`
	IsAvailable  bool
	ContractType string `sql:"type:enum('cdi','cdd','interim');DEFAULT:'cdi'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Costumer   Costumer `gorm:"foreignkey:CustomerID;association_foreignkey:ID"`
	CustomerID int
}
