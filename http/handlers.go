package http

import "net/http"

func IssueAccessToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK!"))
}
