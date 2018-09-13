package models

import (
	"time"

	"github.com/kamalraimi/recruiter-api/config"
)

type Relocation struct {
	ID                    uint   `gorm:"primary_key" form:"id" json:"id"`
	StartDate             string `gorm:"not null" form:"start_date" json:"start_date"`
	EndDate               string `gorm:"not null" form:"end_date" json:"end_date"`
	Accomodated           bool   `sql:"DEFAULT:0" form:"accomodated" json:"accomodated" `
	Address               string `gorm:"type:varchar(255)" form:"address" json:"address"`
	TelephoneSubscriber   bool   `sql:"DEFAULT:0" form:"telephone_subscriber" json:"telephone_subscriber"`
	ElectricitySubscriber bool   `sql:"DEFAULT:0" form:"electricity_subscriber" json:"electricity_subscriber"`
	Status                string `sql:"type:enum('waiting','inProgress','completed','canceled');DEFAULT:'waiting'"`
	CreatedAt             time.Time
	UpdatedAt             time.Time

	Collaborater   Collaborater `gorm:"foreignkey:CollaboraterID;association_foreignkey:ID" json:"collaborater"`
	CollaboraterID int
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
