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
	//Creation des tables avec un nom
	DB.SingularTable(true)
	//Mise a jour automatique du schema
	//Db.AutoMigrate(&models.User, &models.Role)

}

func GetDB() *gorm.DB {
	return DB
}

/*func CreateDb() *gorm.DB {
	// Ouverture du fichier
	db, err := gorm.Open("mysql", "root:Shieva/777@/recruiter?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)

	// Erreur de chargement
	if err != nil {
		panic(err)
	}

	return db
}*/
