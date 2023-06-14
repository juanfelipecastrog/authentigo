package http

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func ValidateExpireDate(claims jwt.MapClaims) error {
	expireDate, err := claims.GetExpirationTime()
	if err != nil {
		return jwt.ErrInvalidKey
	}
	currentTime := time.Now()
	if currentTime.After(expireDate.Time) {
		return jwt.ErrTokenExpired
	}
	return nil
}
