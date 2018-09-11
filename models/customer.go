package models

import (
	"time"

	"github.com/kamalraimi/recruiter-api/config"
)

type Customer struct {
	ID             uint   `gorm:"primary_key"`
	Name           string `gorm:"type:varchar(100);unique_index"`
	Description    string `gorm:"type:text"`
	Representative string
	Location       string
	Email          string `gorm:"type:varchar(100);unique_index"`
	Tel            string
	Fax            string
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
	err := config.GetDB().Create(&customer).Error
	return customer, err

}

func UpdateCustomer(customer *Customer) (*Customer, error) {
	err := config.GetDB().Save(&customer).Error
	return customer, err

}

func DeleteCustomer(id string) error {
	return config.GetDB().Where("id = ?", id).Delete(&Customer{}).Error
}
