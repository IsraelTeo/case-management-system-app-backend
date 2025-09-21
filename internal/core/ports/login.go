package ports

import (
	"context"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/payload"
	"github.com/labstack/echo/v4"
)

// LoginService define los métodos de la lógica de negocio relacionados con el login.
type LoginService interface {
	Login(ctx context.Context, request *payload.LoginRequest) (string, error)
}

// LoginHandler define los métodos para manejar las solicitudes HTTP relacionadas con el login.
type LoginHandler interface {
	Login(c echo.Context) error
}
