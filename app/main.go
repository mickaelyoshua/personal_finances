package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Setup routes
	SetupRoutes(router)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}