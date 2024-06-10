package controllers

import (
	"dbo/models"
	"dbo/services"
	"net/http"
	"strconv"

	i "dbo/interfaces"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service services.OrderService
}

var _ = i.ErrorResponseWrapper{}
var _ = i.SuccessResponseWrapper{}

func NewOrderController(service services.OrderService) *OrderController {
	return &OrderController{service}
}

// GetOrders godoc
// @Summary Get all orders
// @Description Get all orders with pagination
// @Tags orders
// @Produce  json
// @Param   page query int false "Page number"
// @Param   limit query int false "Page size"
// @Success 200 {array} models.Order
// @Failure 400 {object} i.ErrorResponseWrapper
// @Failure 500 {object} i.ErrorResponseWrapper
// @Router /orders [get]
func (ctrl *OrderController) GetAllOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	orders, total, err := ctrl.service.GetAllOrders(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orders, "total": total})
}

// GetOrder godoc
// @Summary Get an order
// @Description Get order by ID
// @Tags orders
// @Produce  json
// @Param   id path int true "Order ID"
// @Success 200 {object} models.Order
// @Failure 400 {object} i.ErrorResponseWrapper
// @Failure 404 {object} i.ErrorResponseWrapper
// @Failure 500 {object} i.ErrorResponseWrapper
// @Router /orders/{id} [get]
func (ctrl *OrderController) GetOrderByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := ctrl.service.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the provided details
// @Tags orders
// @Accept  json
// @Produce  json
// @Param   order body models.Order true "Order"
// @Success 201 {object} models.Order
// @Failure 400 {object} i.ErrorResponseWrapper
// @Failure 500 {object} i.ErrorResponseWrapper
// @Router /orders [post]
func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdOrder, err := ctrl.service.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": createdOrder})
}

// UpdateOrder godoc
// @Summary Update an order
// @Description Update order by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param   id path int true "Order ID"
// @Param   order body models.Order true "Order"
// @Success 200 {object} models.Order
// @Failure 400 {object} i.ErrorResponseWrapper
// @Failure 404 {object} i.ErrorResponseWrapper
// @Failure 500 {object} i.ErrorResponseWrapper
// @Router /orders/{id} [put]
func (ctrl *OrderController) UpdateOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.ID = uint(id)
	updatedOrder, err := ctrl.service.UpdateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updatedOrder})
}

// DeleteOrder godoc
// @Summary Delete an order
// @Description Delete order by ID
// @Tags orders
// @Produce  json
// @Param   id path int true "Order ID"
// @Success 204 {object} i.SuccessResponseWrapper
// @Failure 400 {object} i.ErrorResponseWrapper
// @Failure 404 {object} i.ErrorResponseWrapper
// @Failure 500 {object} i.ErrorResponseWrapper
// @Router /orders/{id} [delete]
func (ctrl *OrderController) DeleteOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.service.DeleteOrder(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Order deleted"})
}
