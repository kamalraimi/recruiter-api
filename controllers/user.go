package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

func GetUsers(c *gin.Context) {
	if users, err := models.FindAllUser(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}
	// curl -i http://localhost:8080/users/
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	if user, err := models.FindUserById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
	// curl -i http://localhost:8080/users/1
}

func PostUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	if user, err := models.CreateUser(&user); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}

	// curl -d '{"name":"admin"}' -H "Content-Type: application/json" -X POST http://localhost:8080/users
}

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

	// curl -d '{"name":"admin_"}' -H "Content-Type: application/json" -X PUT http://localhost:8080/users/1
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeleteUser(id); err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

	// curl -i -X DELETE http://localhost:8080/users/1

}
