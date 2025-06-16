package util

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)


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

func ExecSQLScript(conn *pgx.Conn, scriptPath string) error {
	script, err := os.ReadFile(scriptPath)
	if err != nil {
		return err
	}
	_, err = conn.Exec(context.Background(), string(script))
	return err
}

func GetConn(ctx context.Context, databaseURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return nil, errors.New("failed to connect to database: " + err.Error())
	}
	return conn, nil
}