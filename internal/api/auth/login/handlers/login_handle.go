package handlers

import (
    "net/http"

    "go-links/internal/api/auth/login/models"
	"go-links/internal/entities"

    "github.com/labstack/echo/v5"
)

func (h *AuthHandler) Login(c *echo.Context) error {
    var req models.LoginRequest
    
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
    }

    res, err := h.service.CompareUser(c.Request().Context(), &entities.User{}, &req)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, res)
}