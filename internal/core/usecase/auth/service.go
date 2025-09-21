package auth

import (
	"github.com/IsraelTeo/case-management-system-app-backend/internal/app/config"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/ports"
)

// Service implementa la interfaz LoginService.
type Service struct {
	repository ports.UserRepository
	cfg        *config.JWT
}

// New crea una nueva instancia de Service.
func New(repository ports.UserRepository, cfg *config.JWT) ports.LoginService {
	return &Service{
		repository: repository,
		cfg:        cfg,
	}
}
