package controllers

import (
	"dbo/models"
	"dbo/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	service services.CustomerService
}

func NewCustomerController(service services.CustomerService) *CustomerController {
	return &CustomerController{service}
}

func (ctrl *CustomerController) GetAllCustomers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	customers, total, err := ctrl.service.GetAllCustomers(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customers, "total": total})
}

func (ctrl *CustomerController) GetCustomerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	customer, err := ctrl.service.GetCustomerByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func (ctrl *CustomerController) CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdCustomer, err := ctrl.service.CreateCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": createdCustomer})
}

func (ctrl *CustomerController) UpdateCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer.ID = uint(id)
	updatedCustomer, err := ctrl.service.UpdateCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updatedCustomer})
}

func (ctrl *CustomerController) DeleteCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.service.DeleteCustomer(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Customer deleted"})
}
