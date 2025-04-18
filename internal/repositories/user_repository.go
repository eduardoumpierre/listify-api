package repositories

import (
	"listify-api/internal/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) FindAll() ([]models.User, error) {
	return []models.User{
		{ID: "1", Name: "John Doe"},
		{ID: "2", Name: "Jane Doe"},
	}, nil
}
