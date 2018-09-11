package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/config"
	"github.com/kamalraimi/recruiter-api/models"
	"github.com/kamalraimi/recruiter-api/routing"
)

func main() {

	// Initialisation de la BD et generation des tables
	config.InitDB()
	config.GetDB().AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Collaborater{},
	)

	// Gestion du routage et demarrage du serveur
	router := gin.Default()
	router.RedirectTrailingSlash = true // Support de la barre oblique optionnelle les URLS
	routing.Routes(router)
	router.Run(":8080")

}
