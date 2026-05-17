package handlers

import (
	"net/http"
	"github.com/labstack/echo/v5"
)

func (h *LinkHandler) Redirect(c *echo.Context) error {
	shortCode := c.Param("shortCode")
	
	originalURL, err := h.service.GetOriginalURL(c.Request().Context(), shortCode)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "link not found"})
	}

	return c.Redirect(http.StatusFound, originalURL)
}
