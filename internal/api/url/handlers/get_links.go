package handlers

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

// GetMyLinks returns all links created by the authenticated user
func (h *LinkHandler) GetMyLinks(c *echo.Context) error {
	userID, err := uuid.Parse(c.Get("user_id").(string))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid user_id"})
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit < 1 {
		limit = 10
	}

	res, err := h.service.GetUserLinks(c.Request().Context(), userID, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

// GetLinkByID returns a single link by its UUID
func (h *LinkHandler) GetLinkByID(c *echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid link id"})
	}

	link, err := h.service.GetLinkByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "link not found"})
	}

	return c.JSON(http.StatusOK, link)
}

// GetLinkStats returns click analytics for a specific link
func (h *LinkHandler) GetLinkStats(c *echo.Context) error {
	linkID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid link id"})
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit < 1 {
		limit = 20
	}

	res, err := h.service.GetLinkStats(c.Request().Context(), linkID, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
