package customer

import (
	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/storage/postgres"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/ports"
)

// Repository representa un repositorio para la entidad cliente.
type Repository struct {
	db *postgres.DB
}

// New crea una nueva instancia del repositorio de cliente.
func New(db *postgres.DB) ports.CustomerRepository {
	return &Repository{
		db: db,
	}
}
