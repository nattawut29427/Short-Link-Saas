package handlers

import (
	"fmt"
	"go-links/internal/api/url/models"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h *LinkHandler) CreateUrl(c *echo.Context) error {
	var req models.UrlRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
	}

	baseURL := fmt.Sprintf("%s://%s", c.Scheme(), c.Request().Host)

	res, err := h.service.CreateLink(c.Request().Context(), &req, baseURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}
