package models

import (
	"time"
)

type Menu struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	UrlIcon   string
	FileMenu  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
