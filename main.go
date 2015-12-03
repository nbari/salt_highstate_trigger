package main

import (
	"github.com/nbari/violetear"
	"log"
	"net/http"
)

var version = "0.0.2"

func main() {
	router := violetear.New()
	router.LogRequests = true

	router.HandleFunc("/*", mainHandler)
	router.HandleFunc("/_healthcheck_", healthCheck)

	log.Fatal(http.ListenAndServe(":8080", router))
}
