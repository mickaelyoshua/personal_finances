package util

import (
	"context"
	"errors"
	"os"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func CompareHashedPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}


func GenerateToken(userID int32) (string, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(), // Token valid for 72 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ParseAndValidateToken(tokenString string) (jwt.MapClaims, error) {
	// Retrieve the secret key
	secretKey, err := getSecretKey()
	if err != nil {
		return nil, errors.New("failed to retrieve secret key: " + err.Error())
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, errors.New("failed to parse token: " + err.Error())
	}

	// Validate the token
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to extract claims from token")
	}

	return claims, nil
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