package models

import (
	"time"

	"github.com/kamalraimi/recruiter-api/config"
)

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

	Position Position `gorm:"foreignkey:PostID;association_foreignkey:ID"`
	PostID   int
}

func FindAllApplication() ([]Application, error) {
	var applications []Application
	err := config.GetDB().Find(&applications).Error
	return applications, err
}

func FindApplicationById(id string) (Application, error) {
	var application Application
	err := config.GetDB().Where("id = ?", id).First(&application).Error
	return application, err
}

func CreateApplication(application *Application) (*Application, error) {
	err := config.GetDB().Create(&application).Error
	return application, err

}

func UpdateApplication(application *Application) (*Application, error) {
	err := config.GetDB().Save(&application).Error
	return application, err

}

func DeleteApplication(id string) error {
	return config.GetDB().Where("id = ?", id).Delete(&Application{}).Error
}
