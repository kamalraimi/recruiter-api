package models

import (
	"time"

	"github.com/kamalraimi/recruiter-api/config"
)

type Immigration struct {
	ID        uint `gorm:"primary_key" form:"id" json:"id"`
	StartDate time.Time
	EndDate   time.Time
	Visa      bool
	Status    string `sql:"type:enum('inProgress','completed','canceled');DEFAULT:'inProgress'"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Collaborater   Collaborater `gorm:"foreignkey:CollaboraterID;association_foreignkey:ID"`
	CollaboraterID int
}

func FindAllImmigration() ([]Immigration, error) {
	var immigrations []Immigration
	err := config.GetDB().Find(&immigrations).Error
	return immigrations, err
}

func FindImmigrationById(id string) (Immigration, error) {
	var immigration Immigration
	err := config.GetDB().Where("id = ?", id).First(&immigration).Error
	return immigration, err
}

func CreateImmigration(immigration *Immigration) (*Immigration, error) {
	err := config.GetDB().Create(&immigration).Error
	return immigration, err

}

func UpdateImmigration(immigration *Immigration) (*Immigration, error) {
	err := config.GetDB().Save(&immigration).Error
	return immigration, err

}

func DeleteImmigration(id string) error {
	return config.GetDB().Where("id = ?", id).Delete(&Immigration{}).Error
}
