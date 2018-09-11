package models

import "time"

type Collaborater struct {
	ID        uint   `gorm:"primary_key"`
	Civility  string `sql:"type:enum('mr','mrs','miss');DEFAULT:'mr'"`
	Name      string `gorm:"type:varchar(100);index"`
	BirthDay  time.Time
	Email     string `gorm:"type:varchar(100);unique_index"`
	Tel       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
