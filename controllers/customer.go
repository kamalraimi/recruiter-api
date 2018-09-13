package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamalraimi/recruiter-api/models"
)

// @Summary Consultation des clients
// @Description Recupere les clients
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Customer
// @Router /customers/ [get]
func GetCustomers(c *gin.Context) {
	if customers, err := models.FindAllCustomer(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customers)
	}

}

// @Summary Consultation d'un client
// @Description Recupere un client
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Customer
// @Router /customers/{id} [get]
func GetCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	if customer, err := models.FindCustomerById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customer)
	}

}

// @Summary Enregistrer un client
// @Description Stocke un client
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Customer
// @Router /customers/ [post]
func PostCustomer(c *gin.Context) {
	var customer models.Customer
	c.BindJSON(&customer)
	if customer, err := models.CreateCustomer(&customer); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, customer)
	}

}

// @Summary Modification d'un client
// @Description modifier un client
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Customer
// @Router /customers/{id} [put]
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

}

// @Summary Suppression de client
// @Description supprime un client
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Customer
// @Router /customers/{id} [delete]
func DeleteCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := models.DeleteCustomer(id); err != nil {
		c.AbortWithStatus(405)
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})

}
