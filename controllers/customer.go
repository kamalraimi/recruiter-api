package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

func GetCustomers(c *gin.Context) {
	if customers, err := models.FindAllCustomer(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customers)
	}
	// curl -i http://localhost:8080/customers/
}

func GetCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	if customer, err := models.FindCustomerById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customer)
	}
	// curl -i http://localhost:8080/customers/1
}

func PostCustomer(c *gin.Context) {
	var customer models.Customer
	c.BindJSON(&customer)
	if customer, err := models.CreateCustomer(&customer); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customer)
	}

	// curl -d '{"name":"admin"}' -H "Content-Type: application/json" -X POST http://localhost:8080/customers
}

func PutCustomer(c *gin.Context) {

	id := c.Params.ByName("id")
	customer, err := models.FindCustomerById(id)
	fmt.Println(customer)
	if customer.ID == 0 {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.BindJSON(&customer)
	customerUpdated, err := models.UpdateCustomer(&customer)
	c.JSON(200, &customerUpdated)

	// curl -d '{"name":"admin_"}' -H "Content-Type: application/json" -X PUT http://localhost:8080/customers/1
}

func DeleteCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeleteCustomer(id); err != nil {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

	// curl -i -X DELETE http://localhost:8080/customers/1

}
