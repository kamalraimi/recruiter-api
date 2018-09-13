package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

// @Summary Consultation des comptes
// @Description Recupere les comptes
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Router /users/ [get]
func GetUsers(c *gin.Context) {
	if users, err := models.FindAllUser(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}
}

// @Summary Consultation d'un compte
// @Description Recupere d'un compte
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	if user, err := models.FindUserById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}

// @Summary Enregistrer un compte
// @Description Stocke un compte
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Router /users/ [post]
func PostUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	if user, err := models.CreateUser(&user); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}

}

// @Summary Modification d'un user
// @Description modifier un user
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func PutUser(c *gin.Context) {

	id := c.Params.ByName("id")
	user, err := models.FindUserById(id)
	fmt.Println(user)
	if user.ID == 0 {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.BindJSON(&user)
	userUpdated, err := models.UpdateUser(&user)
	c.JSON(200, &userUpdated)

}

// @Summary Suppression de compte
// @Description supprime une compte
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeleteUser(id); err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
