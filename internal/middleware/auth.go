package middleware

import (
	"log/slog"
	"net/http"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/jwt"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/app/config"
	"github.com/labstack/echo/v4"
)

// ValidateJWT es un middleware que valida el token JWT en las solicitudes entrantes.
// Tambien extrae el userID del token y lo agrega al contexto de la solicitud.
func ValidateJWT(next echo.HandlerFunc, cfg *config.JWT) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := jwt.ExtractUserID(c, cfg)
		if err != nil {
			slog.Error("Error extracting user ID from token", slog.String("error", err.Error()))
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid or missing token",
			})
		}

		return next(c)
	}
}
