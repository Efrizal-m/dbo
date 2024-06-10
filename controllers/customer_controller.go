package controllers

import (
	"dbo/models"
	"dbo/services"
	"net/http"
	"strconv"

	i "dbo/interfaces"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	service services.CustomerService
}

var _ = i.ErrorResponseWrapper{}
var _ = i.SuccessResponseWrapper{}

func NewCustomerController(service services.CustomerService) *CustomerController {
	return &CustomerController{service}
}

// GetCustomers godoc
// @Summary Get all customers
// @Description Get all customers with pagination
// @Tags customers
// @Produce  json
// @Param   page query int false "Page number"
// @Param   limit query int false "Page size"
// @Success 200 {array} models.Customer
// @Failure 400 {object} i.ErrorResponseWrapper
// @Failure 500 {object} i.ErrorResponseWrapper
// @Router /customers [get]
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

// GetCustomer godoc
// @Summary Get a customer
// @Description Get customer by ID
// @Tags customers
// @Produce  json
// @Param   id path int true "Customer ID"
// @Success 200 {object} models.Customer
// @Failure 400 {object} i.ErrorResponseWrapper
// @Failure 404 {object} i.ErrorResponseWrapper
// @Failure 500 {object} i.ErrorResponseWrapper
// @Router /customers/{id} [get]
func (ctrl *CustomerController) GetCustomerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	customer, err := ctrl.service.GetCustomerByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// CreateCustomer godoc
// @Summary Create a new customer
// @Description Create a new customer with the provided details
// @Tags customers
// @Accept  json
// @Produce  json
// @Param   customer body models.Customer true "Customer"
// @Success 201 {object} models.Customer
// @Failure 400 {object} i.ErrorResponseWrapper
// @Failure 500 {object} i.ErrorResponseWrapper
// @Router /customers [post]
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

// UpdateCustomer godoc
// @Summary Update a customer
// @Description Update customer by ID
// @Tags customers
// @Accept  json
// @Produce  json
// @Param   id path int true "Customer ID"
// @Param   customer body models.Customer true "Customer"
// @Success 200 {object} models.Customer
// @Failure 400 {object} i.ErrorResponseWrapper
// @Failure 404 {object} i.ErrorResponseWrapper
// @Failure 500 {object} i.ErrorResponseWrapper
// @Router /customers/{id} [put]
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

// DeleteCustomer godoc
// @Summary Delete a customer
// @Description Delete customer by ID
// @Tags customers
// @Produce  json
// @Param   id path int true "Customer ID"
// @Success 204 {object} i.SuccessResponseWrapper
// @Failure 400 {object} i.ErrorResponseWrapper
// @Failure 404 {object} i.ErrorResponseWrapper
// @Failure 500 {object} i.ErrorResponseWrapper
// @Router /customers/{id} [delete]
func (ctrl *CustomerController) DeleteCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.service.DeleteCustomer(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Customer deleted"})
}
