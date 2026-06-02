package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go-links/internal/api/auth/oauth/models"
	"go-links/internal/api/auth/oauth/repositories"
	"go-links/internal/entities"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuthService interface {
	AuthenticateWithGoogle(ctx context.Context, code string) (*models.OAuthResponse, error)
}

type oauthService struct {
	repo        repositories.OAuthRepository
	jwtSecret   string
	clientID    string
	clientSecret string
	redirectURL string
}

func NewOAuthService(repo repositories.OAuthRepository, jwtSecret, clientID, clientSecret, redirectURL string) OAuthService {
	return &oauthService{
		repo:        repo,
		jwtSecret:   jwtSecret,
		clientID:    clientID,
		clientSecret: clientSecret,
		redirectURL: redirectURL,
	}
}

func (s *oauthService) AuthenticateWithGoogle(ctx context.Context, code string) (*models.OAuthResponse, error) {
	// Get Google OAuth2 configuration from secrets config
	config := &oauth2.Config{
		ClientID:     s.clientID,
		ClientSecret: s.clientSecret,
		RedirectURL:  s.redirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	// Exchange code for token
	token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange token: %w", err)
	}

	// Get user info from Google API
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info: %d", resp.StatusCode)
	}

	// Parse the response and extract user data
	var userData map[string]interface{}
	if err := decodeJSON(resp.Body, &userData); err != nil {
		return nil, fmt.Errorf("failed to decode user data: %w", err)
	}

	email, ok := userData["email"].(string)
	if !ok {
		return nil, fmt.Errorf("email not found in user data")
	}

	name, _ := userData["name"].(string)

	// Create or update user in our database
	user := &entities.User{
		Email: email,
		Title: name, // Using the name as the title
	}
	
	dbUser, err := s.repo.CreateOrUpdateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create or update user: %w", err)
	}

	// Generate JWT token
	tokenString, err := s.generateJWT(*dbUser)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	authResponse := &models.OAuthResponse{
		Token: tokenString,
		User: models.User{
			ID:    dbUser.ID.String(),
			Email: dbUser.Email,
			Name:  dbUser.Title,
		},
	}

	return authResponse, nil
}

func (s *oauthService) generateJWT(user entities.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Simple JSON decoder function since we don't want to import encoding/json in addition to other packages
func decodeJSON(r io.Reader, v interface{}) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}