package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}
func (maker *JWTMaker) CreateToken(userID int32, duration time.Duration) (string, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	// Check the algorithm used in the token
	keyFunc := func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrorExpiredToken
		}
		return nil, ErrorInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok || !jwtToken.Valid {
		return nil, ErrorInvalidToken
	}

	return payload, nil
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid secret key size, must be at least %d characters", minSecretKeySize)
	}

	return &JWTMaker{
		secretKey: secretKey,
	}, nil
}