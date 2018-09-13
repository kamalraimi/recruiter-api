package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

// @Summary Consultation des roles
// @Description Recupere les roles
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Role
// @Router /roles/ [get]
func GetRoles(c *gin.Context) {
	if roles, err := models.FindAllRole(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, roles)
	}

}

// @Summary Consultation d'un rôle
// @Description Recupere d'un rôle
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Role
// @Router /roles/{id} [get]
func GetRole(c *gin.Context) {
	id := c.Params.ByName("id")
	if role, err := models.FindRoleById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, role)
	}

}

// @Summary Enregistrer un role
// @Description Stocke un role
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Role
// @Router /roles/ [post]
func PostRole(c *gin.Context) {
	var role models.Role
	c.BindJSON(&role)
	if role, err := models.CreateRole(&role); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, role)
	}

}

// @Summary Modification d'un role
// @Description modifier un role
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Role
// @Router /roles/{id} [put]
func PutRole(c *gin.Context) {

	id := c.Params.ByName("id")
	role, err := models.FindRoleById(id)
	fmt.Println(role)
	if role.ID == 0 {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.BindJSON(&role)
	roleUpdated, err := models.UpdateRole(&role)
	c.JSON(200, &roleUpdated)

}

// @Summary Suppression de role
// @Description supprime un role
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Role
// @Router /roles/{id} [delete]
func DeleteRole(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeleteRole(id); err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

}
