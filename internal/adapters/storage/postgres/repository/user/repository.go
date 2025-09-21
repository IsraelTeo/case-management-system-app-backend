package user

import (
	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/storage/postgres"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/ports"
)

// Repository representa un repositorio para la entidad usuario.
type Repository struct {
	db *postgres.DB
}

// New crea una nueva instancia del repositorio de usuario.
func New(db *postgres.DB) ports.UserRepository {
	return &Repository{
		db: db,
	}
}
