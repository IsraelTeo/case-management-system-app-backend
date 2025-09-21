package customer

import (
	"net/http"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/jwt"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/payload"
	"github.com/labstack/echo/v4"
)

// Create maneja la creación de un nuevo cliente.
func (h *Handler) Create(c echo.Context) error {
	userID, err := jwt.ExtractUserID(c, h.cfg)
	if err != nil {
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{
				"error": "invalid token",
			},
		)
	}

	customerReq := &payload.CreateCustomerRequest{}
	if err := c.Bind(customerReq); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid request payload",
			},
		)
	}

	createdCustomer, err := h.service.Create(c.Request().Context(), customerReq, userID)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": "failed to create customer",
			},
		)
	}

	return c.JSON(http.StatusCreated, createdCustomer)
}
