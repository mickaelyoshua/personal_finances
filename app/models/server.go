package models

import (
	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua/personal-finances/db/sqlc"
	"github.com/mickaelyoshua/personal-finances/handlers"
	"github.com/mickaelyoshua/personal-finances/middlewares"
)

type Server struct {
	agent  *SQLAgent
	router *gin.Engine
}

func NewServer(agent *SQLAgent) *Server {
	router := gin.Default()

	SetUpRoutes(router, agent.Queries)

	server := &Server{
		agent: agent,
		router: router,
	}
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func SetUpRoutes(router *gin.Engine, queries *sqlc.Queries) {
	// Public routes
	router.GET("/health", handlers.HealthCheck)

	// Auth routes
	authGroup := router.Group("/auth")
	authGroup.POST("/register", handlers.Register(queries))
	authGroup.GET("/register", handlers.RegisterView)
	authGroup.POST("/login", handlers.Login(queries))
	authGroup.GET("/login", handlers.LoginView)

	// Protected routes
	protectedGroup := router.Group("/")
	protectedGroup.Use(middlewares.AuthMiddleware())
	protectedGroup.GET("/", handlers.Index(queries))
	//protectedGroup.GET("/user", handlers.UserView)
	//protectedGroup.POST("/user", handlers.CreateUser)
	//protectedGroup.PUT("/user", handlers.UpdateUser)
	//protectedGroup.DELETE("/user", handlers.DeleteUser)
}