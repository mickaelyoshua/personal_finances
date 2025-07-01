package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua/personal_finances/token"
)

const (
	authorizationHeaderKey = "Authorization"
	authorizationTypeBearer = "Bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader(authorizationHeaderKey)
		if authorizationHeader == "" {
			log.Printf("Authorization header is missing\n")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			// c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			log.Printf("Invalid authorization header format\n")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			// c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		authorizationType := fields[0]
		if authorizationType != authorizationTypeBearer {
			log.Printf("Unsupported authorization type: %s\n", authorizationType)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unsupported authorization type"})
			// c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			log.Printf("Failed to verify token: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			// c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		c.Set(authorizationPayloadKey, payload)
		c.Next()
	}
}