package customer

import (
	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/ports"
)

// Service implementa la interfaz ports.CustomerService.
type Service struct {
	repository ports.CustomerRepository
}

// New crea una nueva instancia del servicio de clientes.
func New(repository ports.CustomerRepository) ports.CustomerService {
	return &Service{
		repository: repository,
	}
}
