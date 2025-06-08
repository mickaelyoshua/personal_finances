package util

import (
	"errors"
	"os"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSecretKey() (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("JWT secret key is not set")
	}
	return secretKey, nil
}

func GetDatabaseURL() (string, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return "", errors.New("DATABASE_URL environment variable is not set")
	}
	return databaseURL, nil
}

func GetTokenFromCookie(c *gin.Context) (string, error) {
	token, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return "", errors.New("no token found in cookie")
		}
		return "", errors.New("failed to retrieve token from cookie: " + err.Error())
	}
	return token, nil
}