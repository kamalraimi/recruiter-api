package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/RichardKnop/go-fixtures"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kamalraimi/recruiter-api/config"
	"github.com/kamalraimi/recruiter-api/models"
	"github.com/kamalraimi/recruiter-api/routing"
)

func main() {

	// Initialisation de la BD et generation des tables
	config.InitDB()
	config.GetDB().Debug().AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Collaborater{},
		&models.Application{},
		&models.Customer{},
		&models.Immigration{},
		&models.Menu{},
		&models.Position{},
		&models.Relocation{},
	)

	//Chargement des données
	loadDataFixtures()

	// Gestion du routage et demarrage du serveur
	router := gin.Default()
	router.RedirectTrailingSlash = true // Support de la barre oblique optionnelle les URLS

	routing.Routes(router)
	router.Run(":8080")

}

func loadDataFixtures() {

	fmt.Println("========================chargement des donnees========================")

	db, err := sql.Open("mysql", config.Connect)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	data, err := ioutil.ReadFile("./fixtures/fixture.yaml")
	print(data)
	if err != nil {
		log.Fatal(err)
	}

	if err := fixtures.Load(data, db, "mysql"); err != nil {
		log.Fatal(err)
	}

}
