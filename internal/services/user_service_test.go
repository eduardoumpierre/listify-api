package services_test

import (
	"context"
	"errors"
	"listify-api/internal/models"
	"listify-api/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockRepository implements repositories.UserRepositoryInterface
type MockRepository struct {
	users []*models.User
	err   error
}

func (m *MockRepository) FindAll(ctx context.Context) ([]*models.User, error) {
	return m.users, m.err
}

func TestFindAllUsers_Success(t *testing.T) {
	mockUsers := []*models.User{
		{ID: "1", Name: "Alice"},
		{ID: "2", Name: "Bob"},
	}

	mockRepo := &MockRepository{users: mockUsers}
	service := services.NewUserService(mockRepo)

	users, err := service.FindAll(context.Background())

	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, "Alice", users[0].Name)
	assert.Equal(t, "Bob", users[1].Name)
}

func TestFindAllUsers_Empty(t *testing.T) {
	mockRepo := &MockRepository{users: []*models.User{}}
	service := services.NewUserService(mockRepo)

	users, err := service.FindAll(context.Background())

	assert.NoError(t, err)
	assert.Len(t, users, 0)
}

func TestFindAllUsers_Error(t *testing.T) {
	mockRepo := &MockRepository{
		users: nil,
		err:   errors.New("database error"),
	}
	service := services.NewUserService(mockRepo)

	users, err := service.FindAll(context.Background())

	assert.Error(t, err)
	assert.Nil(t, users)
}
