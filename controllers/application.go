package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

// @Summary Consultation des candidatures
// @Description Recupere les candidatures
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Application
// @Router /applications/ [get]
func GetApplications(c *gin.Context) {
	if applications, err := models.FindAllApplication(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, applications)
	}

}

// @Summary Consultation d'une candidature donnée
// @Description Recupere une candidature
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Application
// @Router /applications/{id} [get]
func GetApplication(c *gin.Context) {
	id := c.Params.ByName("id")
	if application, err := models.FindApplicationById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, application)
	}

}

// @Summary Enregistrer une candidature
// @Description Stocke une candidature
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Application
// @Router /applications [post]
func PostApplication(c *gin.Context) {
	var application models.Application
	c.BindJSON(&application)
	if application, err := models.CreateApplication(&application); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, application)
	}

}

// @Summary Modification de candidature
// @Description modifier une candidature
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Application
// @Router /applications/{id} [put]
func PutApplication(c *gin.Context) {

	id := c.Params.ByName("id")
	application, err := models.FindApplicationById(id)
	fmt.Println(application)
	if application.ID == 0 {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.BindJSON(&application)
	applicationUpdated, err := models.UpdateApplication(&application)
	c.JSON(200, &applicationUpdated)

}

// @Summary Suppression de candidature
// @Description supprime une candidature
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Application
// @Router /applications/{id} [delete]
func DeleteApplication(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeleteApplication(id); err != nil {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

}
