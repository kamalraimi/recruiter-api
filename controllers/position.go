package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

func GetPositions(c *gin.Context) {
	if positions, err := models.FindAllPosition(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, positions)
	}
	// curl -i http://localhost:8080/positions/
}

func GetPosition(c *gin.Context) {
	id := c.Params.ByName("id")
	if position, err := models.FindPositionById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, position)
	}
	// curl -i http://localhost:8080/positions/1
}

func PostPosition(c *gin.Context) {
	var position models.Position
	c.BindJSON(&position)
	if position, err := models.CreatePosition(&position); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, position)
	}

	// curl -d '{"name":"admin"}' -H "Content-Type: application/json" -X POST http://localhost:8080/positions
}

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

	// curl -d '{"name":"admin_"}' -H "Content-Type: application/json" -X PUT http://localhost:8080/positions/1
}

func DeletePosition(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeletePosition(id); err != nil {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

	// curl -i -X DELETE http://localhost:8080/positions/1

}
