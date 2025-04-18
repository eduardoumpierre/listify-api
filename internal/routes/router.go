package routes

import (
	"listify-api/internal/handlers"
	"listify-api/internal/repositories"
	"listify-api/internal/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go + Air + Docker test!"))
	})

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r.Mount("/users", userHandler.Routes())

	return r
}
