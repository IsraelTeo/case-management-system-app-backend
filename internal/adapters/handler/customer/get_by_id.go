package customer

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetByID maneja la solicitud para obtener un cliente por su ID.
func (h *Handler) GetByID(c echo.Context) error {
	IDStr := c.Param("id")

	ID, err := uuid.Parse(IDStr)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "Invalid ID",
			},
		)
	}

	customer, err := h.service.GetByID(c.Request().Context(), ID)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, customer)
}
