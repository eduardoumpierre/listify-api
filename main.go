package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)

	port := "8080"
	fmt.Println("🚀 Server is running on http://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello from Go + Air + Docker test!"))
}
