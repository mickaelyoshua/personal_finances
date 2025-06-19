package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrorInvalidToken = errors.New("invalid token")
	ErrorExpiredToken = errors.New("token has expired")
)

type Payload struct {
	ID     uuid.UUID `json:"id"`
	UserID int32     `json:"user_id"`
	IssuedAt time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewPayload(userID int32, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:     id,
		UserID: userID,
		IssuedAt: time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (p *Payload) Valid() error {
	if p.ExpiresAt.IsZero() || p.IssuedAt.IsZero() {
		return ErrorInvalidToken
	}

	if time.Now().After(p.ExpiresAt) {
		return ErrorExpiredToken
	}

	if p.ID == uuid.Nil {
		return ErrorInvalidToken
	}

	return nil
}