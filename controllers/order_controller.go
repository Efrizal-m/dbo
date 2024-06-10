package controllers

import (
	"dbo/models"
	"dbo/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service services.OrderService
}

func NewOrderController(service services.OrderService) *OrderController {
	return &OrderController{service}
}

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

func (ctrl *OrderController) GetOrderByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := ctrl.service.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
}

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

func (ctrl *OrderController) DeleteOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.service.DeleteOrder(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Order deleted"})
}
