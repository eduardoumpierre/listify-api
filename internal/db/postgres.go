package db

import (
	"context"
	"fmt"
	"os"

	"listify-api/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return pool, nil
}

func InitializeConfig() (*config.Config, error) {
	pool, err := Connect()
	if err != nil {
		return nil, err
	}

	cfg := &config.Config{
		DB: pool,
	}

	return cfg, nil
}
