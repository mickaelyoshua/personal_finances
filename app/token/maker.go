package token

import "time"

type Maker interface{
	// CreateToken creates a new token for a specific user with a given duration
	CreateToken(userID int32, duration time.Duration) (string, error)

	// VerifyToken checks if the token is valid and returns the payload
	VerifyToken(token string) (*Payload, error)
}