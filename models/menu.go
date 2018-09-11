package models

import (
	"time"
)

type Menu struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(100);unique_index"`
	Type      string `sql:"type:enum('list','detail');DEFAULT:'list'"`
	UrlIcon   string
	FileMenu  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
