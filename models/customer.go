package models

import (
	"time"

	"github.com/kamalraimi/recruiter-api/config"
)

type Customer struct {
	ID             uint   `gorm:"primary_key" form:"id" json:"id"`
	Name           string `gorm:"type:varchar(100);unique_index" form:"name" json:"name"`
	Description    string `gorm:"type:varchar(255);" form:"description" json:"description"`
	Representative string `gorm:"type:varchar(100)" form:"representative" json:"representative"`
	Location       string `gorm:"type:varchar(255)" form:"location" json:"location"`
	Email          string ` gorm:"type:varchar(100);unique_index" form:"email" json:"email"`
	Tel            string `gorm:"type:varchar(25)" form:"tel" json:"tel"`
	Fax            string `gorm:"type:varchar(25)" form:"fax" json:"fax"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func FindAllCustomer() ([]Customer, error) {
	var customers []Customer
	err := config.GetDB().Find(&customers).Error
	return customers, err
}

func FindCustomerById(id string) (Customer, error) {
	var customer Customer
	err := config.GetDB().Where("id = ?", id).First(&customer).Error
	return customer, err
}

func CreateCustomer(customer *Customer) (*Customer, error) {
	err := config.GetDB().Debug().Create(&customer).Error
	return customer, err
}

func UpdateCustomer(customer *Customer) (*Customer, error) {
	err := config.GetDB().Save(&customer).Error
	return customer, err

}

func DeleteCustomer(id string) error {
	return config.GetDB().Where("id = ?", id).Delete(&Customer{}).Error
}
