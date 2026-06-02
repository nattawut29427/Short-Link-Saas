package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"go-links/configs"
	"go-links/internal/api/auth/oauth/models"
	"go-links/internal/api/auth/oauth/services"

	"github.com/labstack/echo/v5"
)

type OAuthHandler struct {
	service services.OAuthService
}

func NewOAuthHandler(service services.OAuthService) *OAuthHandler {
	return &OAuthHandler{
		service: service,
	}
}

func (h *OAuthHandler) GoogleAuthCallback(c *echo.Context) error {
	var req models.GoogleOAuthRequest
	
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
	}

	ctx := context.Background()
	res, err := h.service.AuthenticateWithGoogle(ctx, req.Code)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

// GoogleAuthURL returns the Google OAuth authorization URL
func (h *OAuthHandler) GoogleAuthURL(c *echo.Context) error {
	// Build the Google OAuth URL dynamically based on configurations
	secret := configs.GetSecret()
	clientID := secret.GoogleClientID
	redirectURI := secret.GoogleRedirectURL
	
	if clientID == "" || redirectURI == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Google OAuth is not properly configured",
		})
	}

	authURL := fmt.Sprintf(
		"https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=email%%20profile",
		url.QueryEscape(clientID),
		url.QueryEscape(redirectURI),
	)

	return c.JSON(http.StatusOK, map[string]string{
		"url": authURL,
	})
}