package customer

import (
	"github.com/IsraelTeo/case-management-system-app-backend/internal/app/config"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/ports"
)

// Handler maneja las solicitudes HTTP relacionadas con clientes.
type Handler struct {
	service ports.CustomerService
	cfg     *config.JWT
}

// New crea una nueva instancia de Handler.
func New(service ports.CustomerService, cfg *config.JWT) ports.CustomerHandler {
	return &Handler{
		service: service,
		cfg:     cfg,
	}
}
