package user

import "github.com/IsraelTeo/case-management-system-app-backend/internal/core/ports"

// Handler maneja las solicitudes HTTP relacionadas con usuarios.
type Handler struct {
	service ports.UserService
}
