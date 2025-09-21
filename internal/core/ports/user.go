package ports

import (
	"context"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/domain"
)

// UserRepository define los métodos para acceder y manipular usuarios en la base de datos.
type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
}

// UserService define los métodos de la lógica de negocio relacionados con usuarios.
type UserService interface {
}
