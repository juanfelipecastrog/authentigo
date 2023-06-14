package http

import (
	"Authentigo/auth"
	"Authentigo/local"
	"Authentigo/token"
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
	w.Write([]byte(signedToken))
}
