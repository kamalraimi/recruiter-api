package models

import (
	"time"

	"github.com/kamalraimi/recruiter-api/config"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Login     string `gorm:"type:varchar(50);unique_index"`
	Password  string `gorm:"index"`
	Name      string
	Email     string `gorm:"type:varchar(100);unique_index"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Role   Role `gorm:"foreignkey:RoleID;association_foreignkey:ID"`
	RoleID int
}

func FindAllUser() ([]User, error) {
	var users []User
	err := config.GetDB().Find(&users).Error
	return users, err
}

func FindUserById(id string) (User, error) {
	var user User
	err := config.GetDB().Where("id = ?", id).First(&user).Error
	return user, err
}

func FindUserByName(name string) []User {
	var users []User
	config.GetDB().Find(&users, "name LIKE ?", "%"+name+"%")
	return users

}

func CreateUser(user *User) (uint, error) {
	err := config.GetDB().Create(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil

}

func UpdateUser(user *User) (uint, error) {
	err := config.GetDB().Save(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func DeleteUser(id string) error {
	return config.GetDB().Where("id=?", id).Delete(&Role{}).Error
}
