package customer

import (
	"net/http"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/jwt"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/payload"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Update maneja la solicitud para actualizar un cliente existente.
func (h *Handler) Update(c echo.Context) error {
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

	userID, err := jwt.ExtractUserID(c, h.cfg)
	if err != nil {
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{
				"error": "invalid token",
			},
		)
	}

	customerReq := payload.UpdateCustomerRequest{}
	if err := c.Bind(&customerReq); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid request",
			},
		)
	}

	customerUpdated, err := h.service.Update(c.Request().Context(), &customerReq, ID, userID)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, customerUpdated)
}
