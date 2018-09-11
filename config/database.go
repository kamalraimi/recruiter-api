package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const Connect = "root:Shieva/777@/recruiter?charset=utf8&parseTime=True&loc=Local"

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", Connect)

	if err != nil {
		log.Fatal(err)
	}
	//Creation des tables avec un nom au singulier
	DB.SingularTable(true)

}

func GetDB() *gorm.DB {
	return DB
}
