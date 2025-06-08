package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int32) (string, error) {
	secretKey, err := GetSecretKey()
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
	secretKey, err := GetSecretKey()
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