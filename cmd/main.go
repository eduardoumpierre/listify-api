package main

import (
	"listify-api/internal/db"
	"listify-api/internal/routes"
	"log"
	"net/http"
)

func main() {
	cfg, err := db.InitializeConfig()
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}
	defer cfg.DB.Close()

	r := routes.SetupRouter(cfg)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
