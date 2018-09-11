package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

func GetApplications(c *gin.Context) {
	if applications, err := models.FindAllApplication(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, applications)
	}
	// curl -i http://localhost:8080/applications/
}

func GetApplication(c *gin.Context) {
	id := c.Params.ByName("id")
	if application, err := models.FindApplicationById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, application)
	}
	// curl -i http://localhost:8080/applications/1
}

func PostApplication(c *gin.Context) {
	var application models.Application
	c.BindJSON(&application)
	if application, err := models.CreateApplication(&application); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, application)
	}

	// curl -d '{"name":"admin"}' -H "Content-Type: application/json" -X POST http://localhost:8080/applications
}

func PutApplication(c *gin.Context) {

	id := c.Params.ByName("id")
	application, err := models.FindApplicationById(id)
	fmt.Println(application)
	if application.ID == 0 {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.BindJSON(&application)
	applicationUpdated, err := models.UpdateApplication(&application)
	c.JSON(200, &applicationUpdated)

	// curl -d '{"name":"admin_"}' -H "Content-Type: application/json" -X PUT http://localhost:8080/applications/1
}

func DeleteApplication(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeleteApplication(id); err != nil {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

	// curl -i -X DELETE http://localhost:8080/applications/1

}
