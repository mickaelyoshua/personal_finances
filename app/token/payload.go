package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrorInvalidToken = errors.New("invalid token")
	ErrorExpiredToken = errors.New("token has expired")
)

type Payload struct {
	*jwt.RegisteredClaims // makes the Payload compatible with JWT Claims interface
	ID     uuid.UUID `json:"id"`
	UserID int32     `json:"user_id"`
}

func NewPayload(userID int32, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		RegisteredClaims: &jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
		ID:     id,
		UserID: userID,
	}

	return payload, nil
}