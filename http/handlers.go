package http

import (
	"Authentigo/auth"
	"Authentigo/local"
	"net/http"
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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK!"))
}
