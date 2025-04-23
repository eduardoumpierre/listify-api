package repositories

import (
	"context"
	"listify-api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryInterface interface {
	FindAll(ctx context.Context) ([]*models.User, error)
}

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) FindAll(ctx context.Context) ([]*models.User, error) {
	query := "SELECT id, name FROM users"
	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user = &models.User{}
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
