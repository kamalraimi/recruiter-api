package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

// @Summary Consultation des prj d'immigration
// @Description Recupere les prj d'immigration
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Immigration
// @Router /immigrations/ [get]
func GetImmigrations(c *gin.Context) {
	if immigrations, err := models.FindAllImmigration(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, immigrations)
	}

}

// @Summary Consultation d'un prj d'immigration
// @Description Recupere un prj d'immigration
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Immigration
// @Router /immigrations/{id} [get]
func GetImmigration(c *gin.Context) {
	id := c.Params.ByName("id")
	if immigration, err := models.FindImmigrationById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, immigration)
	}

}

// @Summary Enregistrer un prj d'immigration
// @Description Stocke un prj d'immigration
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Immigration
// @Router /immigrations/ [post]
func PostImmigration(c *gin.Context) {
	var immigration models.Immigration
	c.BindJSON(&immigration)
	if immigration, err := models.CreateImmigration(&immigration); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, immigration)
	}

}

// @Summary Modification d'un prj d'immigration
// @Description modifier un prj d'immigration
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Immigration
// @Router /immigrations/{id} [put]
func PutImmigration(c *gin.Context) {

	id := c.Params.ByName("id")
	immigration, err := models.FindImmigrationById(id)
	fmt.Println(immigration)
	if immigration.ID == 0 {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.BindJSON(&immigration)
	immigrationUpdated, err := models.UpdateImmigration(&immigration)
	c.JSON(200, &immigrationUpdated)

}

// @Summary Suppression d'un prj d'immigration
// @Description supprime un prj d'immigration
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Immigration
// @Router /immigrations/{id} [delete]
func DeleteImmigration(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeleteImmigration(id); err != nil {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

}
