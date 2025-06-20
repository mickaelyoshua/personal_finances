package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua/personal_finances/db/sqlc"
	"github.com/mickaelyoshua/personal_finances/token"
	"github.com/mickaelyoshua/personal_finances/util"
)

type Server struct {
	config    util.Config
	agent     sqlc.Agent
	tokenMaker token.Maker
	router    *gin.Engine
}

func NewServer(config util.Config, agent sqlc.Agent) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create token maker: %w", err)
	}

	router := gin.Default()
	server := &Server{
		config:    config,
		agent:     agent,
		router:    router,
		tokenMaker: tokenMaker,
	}

	server.SetUpRoutes()

	return server, nil
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func (server *Server) SetUpRoutes() {
	// Public routes
	server.router.GET("/health", HealthCheck)

	// Auth routes
	authGroup := server.router.Group("/auth")
	authGroup.POST("/register", server.Register)
	authGroup.GET("/register", RegisterView)
	authGroup.POST("/login", server.Login)
	authGroup.GET("/login", LoginView)

	// Protected routes
	protectedGroup := server.router.Group("/").Use(AuthMiddleware(server.tokenMaker))
	protectedGroup.GET("/", server.Index)
	//protectedGroup.GET("/user", handlers.UserView)
	//protectedGroup.POST("/user", handlers.CreateUser)
	//protectedGroup.PUT("/user", handlers.UpdateUser)
	//protectedGroup.DELETE("/user", handlers.DeleteUser)
}