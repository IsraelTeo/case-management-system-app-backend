package user

import "github.com/IsraelTeo/case-management-system-app-backend/internal/core/ports"

// Service implementa la interfaz ports.UserService.
type Service struct {
	repository ports.UserRepository
}

// New crea una nueva instancia del servicio de usuarios.
func New(repository ports.UserRepository) ports.UserService {
	return &Service{
		repository: repository,
	}
}
