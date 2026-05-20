package services

import (
	"context"

	"go-links/internal/api/auth/login/models"
	"go-links/internal/api/auth/login/repositories"
	"go-links/internal/entities"
)

type LoginService interface {
	// CreateUser(ctx context.Context, req *models.RegisterRequest) (*models.RegisterResponse, error)
	CompareUser(ctx context.Context, user *entities.User, req *models.LoginRequest) (*models.LoginResponse, error)
}

type service struct {
	repo      repositories.LoginRepository
	jwtSecret string
}

func NewAuthService(repo repositories.LoginRepository, jwtSecret string) LoginService {
	return &service{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}
