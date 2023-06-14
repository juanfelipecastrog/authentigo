package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

func GeneratePayload(clientID string, duration time.Duration) jwt.Claims {
	now := time.Now()
	expiration := now.Add(duration)

	return jwt.MapClaims{
		"id":        uuid.New().String(),
		"user":      clientID,
		"issuedAt":  now.Unix(),
		"expiredAt": expiration,
	}
}
