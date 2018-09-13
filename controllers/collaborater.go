package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

// @Summary Consultation des collaborateurs
// @Description Recupere les collaborateurs
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Collaborater
// @Router /collaboraters/ [get]
func GetCollaboraters(c *gin.Context) {
	if collaboraters, err := models.FindAllCollaborater(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, collaboraters)
	}

}

// @Summary Consultation d'un collaborateur
// @Description Recupere un collaborateur
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Collaborater
// @Router /collaborater/{id} [get]
func GetCollaborater(c *gin.Context) {
	id := c.Params.ByName("id")
	if collaborater, err := models.FindCollaboraterById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, collaborater)
	}

}

// @Summary Enregistrer un collaborateur
// @Description Stocke un collaborateur
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Collaborater
// @Router /collaboraters/ [post]
func PostCollaborater(c *gin.Context) {
	var collaborater models.Collaborater
	c.BindJSON(&collaborater)

	if collaborater, err := models.CreateCollaborater(&collaborater); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, collaborater)
	}

}

// @Summary Modification d'un collaborateur
// @Description modifier un collaborateur
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Collaborater
// @Router /collaboraters/{id} [put]
func PutCollaborater(c *gin.Context) {

	id := c.Params.ByName("id")
	collaborater, err := models.FindCollaboraterById(id)
	fmt.Println(collaborater)
	if collaborater.ID == 0 {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.BindJSON(&collaborater)
	collaboraterUpdated, err := models.UpdateCollaborater(&collaborater)
	c.JSON(200, &collaboraterUpdated)

}

// @Summary Suppression d'un collaborateur
// @Description supprime un collaborateur
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Collaborater
// @Router /collaboraters/{id} [delete]
func DeleteCollaborater(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeleteCollaborater(id); err != nil {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

}
