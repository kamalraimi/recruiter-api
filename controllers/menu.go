package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

// @Summary Consultation des menus d'un user
// @Description Utiliser l'id de l'user
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Menu
// @Router /menu/{id} [get]
func GetMenusByUser(c *gin.Context) {
	id := c.Params.ByName("id")
	if menus, err := models.FindAllMenuByUser(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, menus)
	}

}

// @Summary Consultation des menus
// @Description Recupere les menus
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Menu
// @Router /menus/ [get]
func GetMenus(c *gin.Context) {
	if menus, err := models.FindAllMenu(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, menus)
	}

}

// @Summary Consultation d'un menu
// @Description Recupere un menu
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Menu
// @Router /menus/{id} [get]
func GetMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	if menu, err := models.FindMenuById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, menu)
	}

}

// @Summary Enregistrer un menu
// @Description Stocke un menu
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Menu
// @Router /menus/ [post]
func PostMenu(c *gin.Context) {
	var menu models.Menu
	c.BindJSON(&menu)
	if menu, err := models.CreateMenu(&menu); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, menu)
	}

}

// @Summary Modification d'un menu
// @Description modifier un menu
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Menu
// @Router /menus/{id} [put]
func PutMenu(c *gin.Context) {

	id := c.Params.ByName("id")
	menu, err := models.FindMenuById(id)
	if menu.ID == 0 {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.BindJSON(&menu)
	menuUpdated, err := models.UpdateMenu(&menu)
	c.JSON(200, &menuUpdated)

}

// @Summary Suppression de menus
// @Description supprime un menu
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Menu
// @Router /menus/{id} [delete]
func DeleteMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeleteMenu(id); err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

}
