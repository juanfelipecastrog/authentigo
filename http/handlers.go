package http

import (
	"Authentigo/auth"
	"Authentigo/config"
	"Authentigo/local"
	"Authentigo/token"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

// IssueAccessToken issues a JWT access token in response to a request to grant client credentials with basic authentication.
func IssueAccessToken(w http.ResponseWriter, r *http.Request) {
	clientID, clientSecret, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if !auth.VerifyCredentials(clientID, clientSecret, local.Users) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	signedToken, err := token.Create(clientID, time.Duration(24)*time.Hour)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(signedToken))
}

// IntrospectToken verifies and displays information about a JWT access token.
func IntrospectToken(w http.ResponseWriter, r *http.Request) {
	tokenString := r.FormValue("token")
	if tokenString == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if config.PrivateKey == nil || &config.PrivateKey.PublicKey == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	PublicKey := &config.PrivateKey.PublicKey

	appToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return PublicKey, nil
	})

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	claims, ok := appToken.Claims.(jwt.MapClaims)
	if !ok || !appToken.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = ValidateExpireDate(claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	responseJSON, err := json.Marshal(claims)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
