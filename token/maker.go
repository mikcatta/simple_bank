package token

import "time"

type Maker interface {
	// CreateToken creates a new token for a specific username & duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
