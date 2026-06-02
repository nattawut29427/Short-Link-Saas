package handlers

import (
	"net/http"
	"github.com/labstack/echo/v5"
)

func (h *LinkHandler) Redirect(c *echo.Context) error {
	shortCode := c.Param("shortCode")

	ip := c.RealIP()
	userAgent := c.Request().UserAgent()
	referrer := c.Request().Referer()

	originalURL, err := h.service.GetOriginalURLAndRecordClick(c.Request().Context(), shortCode, ip, userAgent, referrer)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "link not found"})
	}

	return c.Redirect(http.StatusFound, originalURL)
}
