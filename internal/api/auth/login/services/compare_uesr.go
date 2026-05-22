package services

import (
	"context"
	"errors"
	"time"

	"go-links/internal/api/auth/login/models"
	"go-links/internal/entities"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

func (s *service) CompareUser(ctx context.Context, user *entities.User, req *models.LoginRequest) (*models.LoginResponse, error) {

	
	user, err := s.repo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if user.Email != req.Email {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return  nil, errors.New("invalid username or password")
	}

	claims := jwt.MapClaims{
        "user_id":  user.ID.String(),
        "email": 	user.Email,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    }
    
    tokenObject := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := tokenObject.SignedString([]byte(s.jwtSecret))
	
    if err != nil {
        return nil, err
    }

	return &models.LoginResponse{
        Token:    tokenString,
    }, nil
}