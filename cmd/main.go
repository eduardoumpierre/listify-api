package main

import (
	"listify-api/internal/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.SetupRouter()

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
