package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

func GetImmigrations(c *gin.Context) {
	if immigrations, err := models.FindAllImmigration(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, immigrations)
	}
	// curl -i http://localhost:8080/immigrations/
}

func GetImmigration(c *gin.Context) {
	id := c.Params.ByName("id")
	if immigration, err := models.FindImmigrationById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, immigration)
	}
	// curl -i http://localhost:8080/immigrations/1
}

func PostImmigration(c *gin.Context) {
	var immigration models.Immigration
	c.BindJSON(&immigration)
	if immigration, err := models.CreateImmigration(&immigration); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, immigration)
	}

	// curl -d '{"name":"admin"}' -H "Content-Type: application/json" -X POST http://localhost:8080/roles
}

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

	// curl -d '{"name":"admin_"}' -H "Content-Type: application/json" -X PUT http://localhost:8080/roles/1
}

func DeleteImmigration(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeleteImmigration(id); err != nil {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

	// curl -i -X DELETE http://localhost:8080/immigrations/1

}
