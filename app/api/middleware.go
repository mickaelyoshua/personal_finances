package api

import (
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
			c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/auth/login")
			return
		}

		c.Set(authorizationPayloadKey, payload)
		c.Next()
	}
}