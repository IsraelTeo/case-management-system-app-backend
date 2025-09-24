package customer

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Delete maneja la solicitud para eliminar un cliente existente.
func (h *Handler) Delete(c echo.Context) error {
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

	if err := h.service.Delete(c.Request().Context(), ID); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK,
		map[string]string{
			"message": "Customer deleted successfully",
		},
	)
}
