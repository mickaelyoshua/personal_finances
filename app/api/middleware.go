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
			c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			log.Printf("Invalid authorization header format\n")
			c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		authorizationType := fields[0]
		if authorizationType != authorizationTypeBearer {
			log.Printf("Unsupported authorization type: %s\n", authorizationType)
			c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			log.Printf("Failed to verify token: %v\n", err)
			c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		c.Set(authorizationPayloadKey, payload)
		c.Next()
	}
}