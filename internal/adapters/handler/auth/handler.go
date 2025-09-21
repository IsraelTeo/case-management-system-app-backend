package auth

import (
	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/ports"
)

// Handler maneja las solicitudes HTTP relacionadas con la autenticación.
type Handler struct {
	service ports.LoginService
}

// New crea una nueva instancia de Handler.
func New(service ports.LoginService) ports.LoginHandler {
	return &Handler{
		service: service,
	}
}
