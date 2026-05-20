package handlers

import (
	"go-links/internal/api/auth/register/models"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h *AuthHandler) CreateUser(c *echo.Context) error {
	var req models.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
	}

	res, err := h.service.CreateUser(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}
