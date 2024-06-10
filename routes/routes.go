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
	// Authentication routes
	userRepository := repositories.NewUserRepository(config.DB)
	authService := services.NewAuthService(userRepository)
	authController := controllers.NewAuthController(authService)

	router.POST("/auth/register", authController.Register)
	router.POST("/auth/login", authController.Login)

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

	// Order routes
	orderService := services.NewOrderService(repositories.NewOrderRepository(config.DB))
	orderController := controllers.NewOrderController(orderService)

	orderRoutes := router.Group("/orders")
	orderRoutes.Use(middlewares.JWTAuthMiddleware())
	{
		orderRoutes.GET("/", orderController.GetAllOrders)
		orderRoutes.GET("/:id", orderController.GetOrderByID)
		orderRoutes.POST("/", orderController.CreateOrder)
		orderRoutes.PUT("/:id", orderController.UpdateOrder)
		orderRoutes.DELETE("/:id", orderController.DeleteOrder)
	}
}
