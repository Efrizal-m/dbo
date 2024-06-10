package main

import (
	"dbo/config"
	"dbo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.ConnectDatabase()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
