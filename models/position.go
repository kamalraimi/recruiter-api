package models

import (
	"time"

	"github.com/kamalraimi/recruiter-api/config"
)

type Position struct {
	ID           uint   `gorm:"primary_key" form:"id" json:"id"`
	Title        string `gorm:"type:varchar(100);index"`
	Description  string `gorm:"type:text"`
	Location     string `gorm:"type:varchar(255)"`
	IsAvailable  bool
	ContractType string `sql:"type:enum('cdi','cdd','interim');DEFAULT:'cdi'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Customer   Customer `gorm:"foreignkey:CustomerID;association_foreignkey:ID"`
	CustomerID int
}

func FindAllPosition() ([]Position, error) {
	var positions []Position
	err := config.GetDB().Find(&positions).Error
	return positions, err
}

func FindPositionById(id string) (Position, error) {
	var position Position
	err := config.GetDB().Where("id = ?", id).First(&position).Error
	return position, err
}

func CreatePosition(position *Position) (*Position, error) {
	err := config.GetDB().Create(&position).Error
	return position, err

}

func UpdatePosition(position *Position) (*Position, error) {
	err := config.GetDB().Save(&position).Error
	return position, err

}

func DeletePosition(id string) error {
	return config.GetDB().Where("id = ?", id).Delete(&Position{}).Error
}
