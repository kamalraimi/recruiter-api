package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

func GetRoles(c *gin.Context) {
	if roles, err := models.FindAllRole(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, roles)
	}
	// curl -i http://localhost:8080/roles/
}

func GetRole(c *gin.Context) {
	id := c.Params.ByName("id")
	if role, err := models.FindRoleById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, role)
	}
	// curl -i http://localhost:8080/roles/1
}

func PostRole(c *gin.Context) {
	var role models.Role
	c.BindJSON(&role)
	if role, err := models.CreateRole(&role); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, role)
	}

	// curl -d '{"name":"admin"}' -H "Content-Type: application/json" -X POST http://localhost:8080/roles
}

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

	// curl -d '{"name":"admin_"}' -H "Content-Type: application/json" -X PUT http://localhost:8080/roles/1
}

func DeleteRole(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeleteRole(id); err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

	// curl -i -X DELETE http://localhost:8080/roles/1

}
