package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua/personal_finances/db/sqlc"
	"github.com/mickaelyoshua/personal_finances/middlewares"
)

type Server struct {
	agent  sqlc.Agent
	router *gin.Engine
}

func NewServer(agent sqlc.Agent) *Server {
	router := gin.Default()

	server := &Server{
		agent: agent,
		router: router,
	}

	SetUpRoutes(server)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func SetUpRoutes(server *Server) {
	// Public routes
	server.router.GET("/health", HealthCheck)

	// Auth routes
	authGroup := server.router.Group("/auth")
	authGroup.POST("/register", server.Register)
	authGroup.GET("/register", RegisterView)
	authGroup.POST("/login", server.Login)
	authGroup.GET("/login", LoginView)

	// Protected routes
	protectedGroup := server.router.Group("/")
	protectedGroup.Use(middlewares.AuthMiddleware())
	protectedGroup.GET("/", server.Index)
	//protectedGroup.GET("/user", handlers.UserView)
	//protectedGroup.POST("/user", handlers.CreateUser)
	//protectedGroup.PUT("/user", handlers.UpdateUser)
	//protectedGroup.DELETE("/user", handlers.DeleteUser)
}