package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

func GenerateClaims(clientID string, duration time.Duration) jwt.Claims {
	return jwt.MapClaims{
		"id":       uuid.New(),
		"user":     clientID,
		"issuedAt": time.Now(),
		"exp":      time.Now().Add(duration).Unix(),
	}
}
