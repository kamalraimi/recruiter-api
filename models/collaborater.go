package models

import (
	"time"

	"github.com/kamalraimi/recruiter-api/config"
)

type Collaborater struct {
	ID        uint   `gorm:"primary_key"`
	Civility  string `sql:"type:enum('mr','mrs','miss');DEFAULT:'mr'"`
	Name      string `gorm:"type:varchar(100);index"`
	BirthDay  time.Time
	Email     string `gorm:"type:varchar(100);unique_index"`
	Tel       string
	CreatedAt time.Time
	UpdatedAt time.Time

	Application   Application `gorm:"foreignkey:ApplicationID;association_foreignkey:ID"`
	ApplicationID int

	Immigration   Immigration `gorm:"foreignkey:ImmigrationID;association_foreignkey:ID"`
	ImmigrationID int

	Relocation   Relocation `gorm:"foreignkey:RelocationID;association_foreignkey:ID"`
	RelocationID int
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
	err := config.GetDB().Create(&collaborater).Error
	return collaborater, err

}

func UpdateCollaborater(collaborater *Collaborater) (*Collaborater, error) {
	err := config.GetDB().Save(&collaborater).Error
	return collaborater, err

}

func DeleteCollaborater(id string) error {
	return config.GetDB().Where("id = ?", id).Delete(&Collaborater{}).Error
}
