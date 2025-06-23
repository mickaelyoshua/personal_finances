package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua/personal_finances/token"
)

const (
	authorizationHeaderKey = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader(authorizationHeaderKey)
		if authorizationHeader == "" {
			fmt.Printf("Authorization header is missing\n")
			c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			fmt.Printf("Invalid authorization header format\n")
			c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			fmt.Printf("Unsupported authorization type: %s\n", authorizationType)
			c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			fmt.Printf("Failed to verify token: %v\n", err)
			c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		c.Set(authorizationPayloadKey, payload)
		c.Next()
	}
}