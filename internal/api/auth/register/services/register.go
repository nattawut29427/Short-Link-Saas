package services

import (
	"context"
	"time"

	"go-links/internal/api/auth/register/models"
	"go-links/internal/entities"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
)

func (s *service) CreateUser(ctx context.Context, req *models.RegisterRequest) (*models.RegisterResponse, error) {

	var user entities.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.ID = uuid.New()
	user.Email = req.Email
	user.Password = string(hashedPassword)

	if err := s.repo.CreateUser(ctx, &user); err != nil {
		return nil, err
	}

	claims := jwt.MapClaims {
		"user_id":  user.ID.String(),                 
        "email":    user.Email, 
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	tokenObject := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := tokenObject.SignedString([]byte(s.jwtSecret))
    if err != nil {
        return nil, err 
    }

	return &models.RegisterResponse{
		ID:       user.ID,
		Email:    user.Email,
		Token: tokenString,

	}, nil
}
