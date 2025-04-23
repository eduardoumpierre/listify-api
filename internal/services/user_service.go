package services

import (
	"context"
	"listify-api/internal/models"
	"listify-api/internal/repositories"
)

type UserService struct {
	repo repositories.UserRepositoryInterface
}

func NewUserService(repo repositories.UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) FindAll(ctx context.Context) ([]*models.User, error) {
	return s.repo.FindAll(ctx)
}
