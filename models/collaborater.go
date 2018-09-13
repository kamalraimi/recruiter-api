package models

import (
	"time"

	"github.com/kamalraimi/recruiter-api/config"
)

type Collaborater struct {
	ID        uint   `gorm:"primary_key"`
	Civility  string `sql:"type:enum('mr','mrs','miss');DEFAULT:'mr'"`
	Name      string `gorm:"type:varchar(100);index" form:"name" json:"name"`
	Email     string `gorm:"type:varchar(100);unique_index" form:"email" json:"email"`
	Tel       string `form:"tel" json:"tel" form:"tel" json:"tel"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Application   Application `gorm:"PRELOAD:false; foreignkey:ApplicationID;association_foreignkey:ID" `
	ApplicationID int
}

func FindAllCollaborater() ([]Collaborater, error) {
	var collaboraters []Collaborater
	err := config.GetDB().Find(&collaboraters).Error
	return collaboraters, err
}

func FindCollaboraterById(id string) (Collaborater, error) {
	var collaborater Collaborater
	err := config.GetDB().Where("id = ?", id).First(&collaborater).Error
	return collaborater, err
}

func CreateCollaborater(collaborater *Collaborater) (*Collaborater, error) {
	err := config.GetDB().Debug().Create(&collaborater).Error
	return collaborater, err

}

func UpdateCollaborater(collaborater *Collaborater) (*Collaborater, error) {
	err := config.GetDB().Save(&collaborater).Error
	return collaborater, err

}

func DeleteCollaborater(id string) error {
	return config.GetDB().Where("id = ?", id).Delete(&Collaborater{}).Error
}
