package auth

import (
	"net/http"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/payload"
	"github.com/labstack/echo/v4"
)

// Login maneja la solicitud de inicio de sesión.
func (h *Handler) Login(c echo.Context) error {
	request := payload.LoginRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "Invalid request",
			},
		)
	}

	token, err := h.service.Login(c.Request().Context(), &request)
	if err != nil {
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusCreated,
		map[string]string{
			"access_token": token,
		},
	)
}
