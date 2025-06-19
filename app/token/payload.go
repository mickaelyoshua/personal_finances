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
	ID        uuid.UUID `json:"id"`
	UserID    int32     `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	jwt.RegisteredClaims // makes the Payload compatible with JWT claims interface
}

func NewPayload(userID int32, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        id,
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}