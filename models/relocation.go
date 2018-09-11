package models

import (
	"time"

	"github.com/kamalraimi/recruiter-api/config"
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

func FindAllRelocation() ([]Relocation, error) {
	var relocations []Relocation
	err := config.GetDB().Find(&relocations).Error
	return relocations, err
}

func FindRelocationById(id string) (Relocation, error) {
	var relocation Relocation
	err := config.GetDB().Where("id = ?", id).First(&relocation).Error
	return relocation, err
}

func CreateRelocation(relocation *Relocation) (*Relocation, error) {
	err := config.GetDB().Create(&relocation).Error
	return relocation, err

}

func UpdateRelocation(relocation *Relocation) (*Relocation, error) {
	err := config.GetDB().Save(&relocation).Error
	return relocation, err

}

func DeleteRelocation(id string) error {
	return config.GetDB().Where("id = ?", id).Delete(&Relocation{}).Error
}
