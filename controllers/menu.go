package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

func GetMenus(c *gin.Context) {
	if menus, err := models.FindAllMenu(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, menus)
	}
	// curl -i http://localhost:8080/menus/
}

func GetMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	if menu, err := models.FindMenuById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, menu)
	}
	// curl -i http://localhost:8080/recruiter-api/menus/1
}

func PostMenu(c *gin.Context) {
	var menu models.Menu
	c.BindJSON(&menu)
	if menu, err := models.CreateMenu(&menu); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, menu)
	}

	// curl -d '{"name":"admin"}' -H "Content-Type: application/json" -X POST http://localhost:8080/menus
}

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

	// curl -d '{"name":"admin_"}' -H "Content-Type: application/json" -X PUT http://localhost:8080/menus/1
}

func DeleteMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeleteMenu(id); err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

	// curl -i -X DELETE http://localhost:8080/menus/1

}
