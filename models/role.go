package models

import (
	"time"

	"github.com/kamalraimi/recruiter-api/config"
)

type Role struct {
	ID        uint   `gorm:"primary_key" form:"id" json:"id"`
	Name      string `gorm:"not null" form:"name" json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FindAllRole() ([]Role, error) {
	var roles []Role
	err := config.GetDB().Find(&roles).Error
	return roles, err
}

func FindRoleById(id string) (Role, error) {
	var role Role
	err := config.GetDB().Where("id = ?", id).First(&role).Error
	return role, err
}

func CreateRole(role *Role) (*Role, error) {
	err := config.GetDB().Create(&role).Error
	return role, err

}

func UpdateRole(role *Role) (*Role, error) {
	err := config.GetDB().Save(&role).Error
	return role, err

}

func DeleteRole(id string) error {
	return config.GetDB().Where("id = ?", id).Delete(&Role{}).Error
}
