package models

import (
	"time"
)

type Relocation struct {
	ID        uint `gorm:"primary_key" form:"id" json:"id"`
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
