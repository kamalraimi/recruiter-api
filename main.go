package main

import (
	"database/sql"
	"io/ioutil"
	"log"

	"github.com/RichardKnop/go-fixtures"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kamalraimi/recruiter-api/config"
	_ "github.com/kamalraimi/recruiter-api/docs"
	"github.com/kamalraimi/recruiter-api/models"
	"github.com/kamalraimi/recruiter-api/routing"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Recruiter API
// @version 1.0
// @description Recruiter-api est une API RESTFUL de gestion des recrutements.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email kamalraimi@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /recruiter-api

// @securityDefinitions.basic BasicAuth
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

	//Chargement des données fixtures
	loadDataFixtures()

	// Gestion du routage et demarrage du serveur
	router := gin.Default()
	router.RedirectTrailingSlash = true // Support de la barre oblique optionnelle les URLS

	//Gestion de la documentation avec swagger
	router.GET("/recruiter-api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routing.Routes(router)
	router.Run(":8080")

}

func loadDataFixtures() {

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
