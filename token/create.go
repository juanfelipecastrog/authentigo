package token

import (
	"Authentigo/config"
	"Authentigo/internal"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func Create(clientID string, duration time.Duration) (string, error) {
	payload := GeneratePayload(clientID, duration)

	privateKey, err := config.ReadPrivateKey(internal.PrivateKeyPath)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, payload)
	tokenString, err := jwtToken.SignedString(privateKey)
	if err != nil {
		fmt.Errorf("an error occurred while signing an API token during creation")
		return "", err
	}

	return tokenString, nil
}
