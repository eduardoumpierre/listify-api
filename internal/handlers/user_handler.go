package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"listify-api/internal/models"
	"listify-api/internal/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.ListUsers)

	return r
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.FindAll(r.Context())
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	if users == nil {
		users = []*models.User{}
	}

	response := struct {
		Users []*models.User `json:"users"`
	}{
		Users: users,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
