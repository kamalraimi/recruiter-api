package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

// @Summary Consultation des relocations
// @Description Recupere les relocations
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Relocation
// @Router /relocations/ [get]
func GetRelocations(c *gin.Context) {
	if relocations, err := models.FindAllRelocation(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, relocations)
	}

}

// @Summary Consultation d'un prj de reloc
// @Description Recupere d'un prj de reloc
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Relocation
// @Router /relocations/{id} [get]
func GetRelocation(c *gin.Context) {
	id := c.Params.ByName("id")
	if relocation, err := models.FindRelocationById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, relocation)
	}

}

// @Summary Enregistrer un prj de reloc
// @Description Stocke un prj de reloc
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Relocation
// @Router /relocations/ [post]
func PostRelocation(c *gin.Context) {
	var relocation models.Relocation
	c.BindJSON(&relocation)
	if relocation, err := models.CreateRelocation(&relocation); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, relocation)
	}

}

// @Summary Modification d'un prj de reloc
// @Description modifier un prj de reloc
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Relocation
// @Router /relocations/{id} [put]
func PutRelocation(c *gin.Context) {

	id := c.Params.ByName("id")
	relocation, err := models.FindRelocationById(id)
	fmt.Println(relocation)
	if relocation.ID == 0 {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.BindJSON(&relocation)
	relocationUpdated, err := models.UpdateRelocation(&relocation)
	c.JSON(200, &relocationUpdated)

}

// @Summary Suppression d'un prj de reloc
// @Description supprime d'un prj de reloc
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Relocation
// @Router /relocations/{id} [delete]
func DeleteRelocation(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeleteRelocation(id); err != nil {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

}
