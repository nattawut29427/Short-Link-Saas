package services

import (
	"context"
	"go-links/internal/api/auth/register/models"
	"go-links/internal/api/auth/register/repositories"

)

type AuthService interface {
	CreateUser(ctx context.Context, req *models.RegisterRequest) (*models.RegisterResponse, error)
}

type service struct {
	repo      repositories.AuthRepository
	jwtSecret string
}

func NewAuthService(repo repositories.AuthRepository, jwtSecret string) AuthService {
	return &service{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}
	