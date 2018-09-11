package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

func GetCollaboraters(c *gin.Context) {
	if collaboraters, err := models.FindAllCollaborater(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, collaboraters)
	}
	// curl -i http://localhost:8080/collaboraters/
}

func GetCollaborater(c *gin.Context) {
	id := c.Params.ByName("id")
	if collaborater, err := models.FindCollaboraterById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, collaborater)
	}
	// curl -i http://localhost:8080/collaboraters/1
}

func PostCollaborater(c *gin.Context) {
	var collaborater models.Collaborater
	c.BindJSON(&collaborater)
	if collaborater, err := models.CreateCollaborater(&collaborater); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, collaborater)
	}

	// curl -d '{"name":"admin"}' -H "Content-Type: application/json" -X POST http://localhost:8080/roles
}

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

	// curl -d '{"name":"admin_"}' -H "Content-Type: application/json" -X PUT http://localhost:8080/roles/1
}

func DeleteCollaborater(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeleteCollaborater(id); err != nil {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

	// curl -i -X DELETE http://localhost:8080/collaboraters/1

}
