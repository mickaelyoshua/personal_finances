package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua/personal_finances/db/sqlc"
	"github.com/mickaelyoshua/personal_finances/token"
	"github.com/mickaelyoshua/personal_finances/util"
)

type Server struct {
	Config    util.Config
	Agent     sqlc.Agent
	TokenMaker token.Maker
	Router    *gin.Engine
}

func NewServer(config util.Config, agent sqlc.Agent) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create token maker: %w", err)
	}

	router := gin.Default()
	server := &Server{
		Config:    config,
		Agent:     agent,
		Router:    router,
		TokenMaker: tokenMaker,
	}

	server.SetUpRoutes()

	return server, nil
}

func (s *Server) Start(address string) error {
	return s.Router.Run(address)
}

func (server *Server) SetUpRoutes() {
	// Public routes
	server.Router.GET("/health", HealthCheck)

	// Auth routes
	authGroup := server.Router.Group("/auth")
	authGroup.POST("/register", server.Register)
	authGroup.GET("/register", RegisterView)
	authGroup.POST("/login", server.Login)
	authGroup.GET("/login", LoginView)

	// Validation routes
	//validateGroup := server.Router.Group("/validate")
	//validateGroup.POST("/email", server.ValidateEmail)

	// Protected routes
	protectedGroup := server.Router.Group("/").Use(AuthMiddleware(server.TokenMaker))
	protectedGroup.GET("/", server.Index)
	//protectedGroup.GET("/user", handlers.UserView)
	//protectedGroup.POST("/user", handlers.CreateUser)
	//protectedGroup.PUT("/user", handlers.UpdateUser)
	//protectedGroup.DELETE("/user", handlers.DeleteUser)
}

func (server *Server) SetToken(c *gin.Context, userID int32) {
	token, err := server.TokenMaker.CreateToken(userID, server.Config.AccessTokenDuration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create access token - " + err.Error()})
		return
	}

	// Set the token in the response header and expose it to the client
	c.Header(authorizationHeaderKey, "Bearer "+token)
	c.Header("Access-Control-Expose-Headers", authorizationHeaderKey)

	// Set the token in a cookie
	c.SetCookie(
		"access_token",
		token,
		int(server.Config.AccessTokenDuration.Seconds()),
		"/",
		"",
		false,
		true,
	)
}