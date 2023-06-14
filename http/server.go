package http

import (
	"log"
	"net/http"
)

func StartServer() {
	router := defineRouter()
	log.Println("authentigo is running...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func defineRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(CreateTokenEndpoint, IssueAccessToken)
	mux.HandleFunc(TokenValidationEndpoint, IntrospectToken)
	mux.HandleFunc(GetKeysEndpoint, ListSigningKeys)
	return mux
}
