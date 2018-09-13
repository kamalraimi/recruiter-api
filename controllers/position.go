package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

// @Summary Consultation des offres d'emploi
// @Description Recupere des offres d'emploi
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Position
// @Router /positions/ [get]
func GetPositions(c *gin.Context) {
	if positions, err := models.FindAllPosition(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, positions)
	}

}

// @Summary Consultation d'une offre d'emploi
// @Description Recupere d'une offre d'emploi
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Menu
// @Router /positions/{id} [get]
func GetPosition(c *gin.Context) {
	id := c.Params.ByName("id")
	if position, err := models.FindPositionById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, position)
	}

}

// @Summary Enregistrer une offre d'emploi
// @Description Stocke une offre d'emploi
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Position
// @Router /positions/ [post]
func PostPosition(c *gin.Context) {
	var position models.Position
	c.BindJSON(&position)
	if position, err := models.CreatePosition(&position); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, position)
	}

}

// @Summary Modification d'une offre d'emploi
// @Description modifier une offre d'emploi
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Position
// @Router /positions/{id} [put]
func PutPosition(c *gin.Context) {

	id := c.Params.ByName("id")
	position, err := models.FindPositionById(id)
	fmt.Println(position)
	if position.ID == 0 {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.BindJSON(&position)
	positionUpdated, err := models.UpdatePosition(&position)
	c.JSON(200, &positionUpdated)

}

// @Summary Suppression d'une offre d'emploi
// @Description supprime une offre d'emploi
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Position
// @Router /positions/{id} [delete]
func DeletePosition(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeletePosition(id); err != nil {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

}
