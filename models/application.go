package models

import "time"

type Application struct {
	ID             uint   `gorm:"primary_key" form:"id" json:"id"`
	Civility       string `sql:"type:enum('mr','mrs','miss');DEFAULT:'mr'"`
	Name           string `gorm:"type:varchar(100);index"`
	BirthDay       time.Time
	Email          string `gorm:"type:varchar(100);index"`
	Tel            string
	ExpectedSalary float64
	Step           string `sql:"type:enum('initialSelection','firstInterview','secondInterview','contractProposal');DEFAULT:'initialSelection'"`
	Status         string `sql:"type:enum('inProgress','validated','rejected');DEFAULT:'inProgress'"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	Position   Position `gorm:"foreignkey:PositionID;association_foreignkey:ID"`
	PositionID int
}
