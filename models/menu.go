package models

import (
	"fmt"
	"time"

	"github.com/kamalraimi/recruiter-api/config"
)

type Menu struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(100);unique_index"`
	Type      string `sql:"type:enum('list','detail');DEFAULT:'list'"`
	UrlIcon   string `gorm:"not null" form:"url_icon" json:"url_icon"`
	FileMenu  string `gorm:"not null" form:"file_menu" json:"file_menu"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FindAllMenu() ([]Menu, error) {
	var menus []Menu
	err := config.GetDB().Find(&menus).Error
	return menus, err
}

func FindAllMenuByUser(id string) ([]Menu, error) {
	var menus []Menu
	role, err := FindRoleById(id)
	if err == nil {
		err = config.GetDB().Joins("JOIN role_menus on role_menus.menu_id=menu.id").
			Where("role_menus.role_id=?", role.ID).Find(&menus).Error
		fmt.Println(menus)
	}
	return menus, err
}

func FindMenuById(id string) (Menu, error) {
	var menu Menu
	err := config.GetDB().Where("id = ?", id).First(&menu).Error
	return menu, err
}

func CreateMenu(menu *Menu) (*Menu, error) {
	err := config.GetDB().Create(&menu).Error
	return menu, err

}

func UpdateMenu(menu *Menu) (*Menu, error) {
	err := config.GetDB().Save(&menu).Error
	return menu, err

}

func DeleteMenu(id string) error {
	return config.GetDB().Where("id = ?", id).Delete(&Menu{}).Error
}
