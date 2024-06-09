package routes

import (
	"dbo/config"
	"dbo/controllers"
	"dbo/middlewares"
	"dbo/repositories"
	"dbo/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Customer routes
	customerService := services.NewCustomerService(repositories.NewCustomerRepository(config.DB))
	customerController := controllers.NewCustomerController(customerService)

	customerRoutes := router.Group("/customers")
	customerRoutes.Use(middlewares.JWTAuthMiddleware())
	{
		customerRoutes.GET("/", customerController.GetAllCustomers)
		customerRoutes.GET("/:id", customerController.GetCustomerByID)
		customerRoutes.POST("/", customerController.CreateCustomer)
		customerRoutes.PUT("/:id", customerController.UpdateCustomer)
		customerRoutes.DELETE("/:id", customerController.DeleteCustomer)
	}
}
