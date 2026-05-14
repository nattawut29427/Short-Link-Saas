package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetAddresses(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := h.service.GetAllAddresses(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *handler) GetAddressByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID format"})
	}

	ctx := c.Request().Context()
	res, err := h.service.GetAddressByID(ctx, uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "address not found"})
	}

	return c.JSON(http.StatusOK, res)
}
