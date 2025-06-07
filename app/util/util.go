package util

import (
	"context"
	"errors"
	"os"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/mickaelyoshua/personal-finances/db/sqlc"
)

type SQLAgent struct {
	Conn    *pgx.Conn
	Queries *sqlc.Queries
}

func getSecretKey() (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("JWT secret key is not set")
	}
	return secretKey, nil
}

func getDatabaseURL() (string, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return "", errors.New("DATABASE_URL environment variable is not set")
	}
	return databaseURL, nil
}

func GetSQLAgent(ctx context.Context) (*SQLAgent, error) {
	databaseURL, err := getDatabaseURL()
	if err != nil {
		return nil, err
	}

	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return nil, err
	}

	queries := sqlc.New(conn)

	return &SQLAgent{
		Conn:    conn,
		Queries: queries,
	}, nil
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