package main

import (
	"Authentigo/http"
	"Authentigo/internal"
	"flag"
	"log"
)

func init() {
	customPrivateKeyPath := flag.String("privateKeyPath", "", "private key file path")
	flag.Parse()
	if *customPrivateKeyPath == "" {
		log.Println("the -privateKeyPath has not been specified, the default private key path will be used")
	} else {
		internal.PrivateKeyPath = *customPrivateKeyPath
	}
}

func main() {
	http.StartServer()
}
