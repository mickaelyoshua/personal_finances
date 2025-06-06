package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua/personal-finances/handlers"
	"github.com/mickaelyoshua/personal-finances/middlewares"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes
	router.GET("/health", handlers.HealthCheck)

	// Auth routes
	authGroup := router.Group("/auth")
	authGroup.POST("/register", handlers.Register)
	authGroup.GET("/register", handlers.RegisterView)
	authGroup.POST("/login", handlers.Login)
	authGroup.GET("/login", handlers.LoginView)

	// Protected routes
	protectedGroup := router.Group("/")
	protectedGroup.Use(middlewares.AuthMiddleware())
	protectedGroup.GET("/", handlers.Index)
	//protectedGroup.GET("/user", handlers.UserView)
	//protectedGroup.POST("/user", handlers.CreateUser)
	//protectedGroup.PUT("/user", handlers.UpdateUser)
	//protectedGroup.DELETE("/user", handlers.DeleteUser)
}

