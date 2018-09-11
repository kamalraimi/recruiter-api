package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

func GetRelocations(c *gin.Context) {
	if relocations, err := models.FindAllRelocation(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, relocations)
	}
	// curl -i http://localhost:8080/relocations/
}

func GetRelocation(c *gin.Context) {
	id := c.Params.ByName("id")
	if relocation, err := models.FindRelocationById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, relocation)
	}
	// curl -i http://localhost:8080/relocations/1
}

func PostRelocation(c *gin.Context) {
	var relocation models.Relocation
	c.BindJSON(&relocation)
	if relocation, err := models.CreateRelocation(&relocation); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, relocation)
	}

	// curl -d '{"name":"admin"}' -H "Content-Type: application/json" -X POST http://localhost:8080/relocations
}

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

	// curl -d '{"name":"admin_"}' -H "Content-Type: application/json" -X PUT http://localhost:8080/relocations/1
}

func DeleteRelocation(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeleteRelocation(id); err != nil {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

	// curl -i -X DELETE http://localhost:8080/relocations/1

}
