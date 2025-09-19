package domain

import (
	"time"

	"github.com/google/uuid"
)

// Headquarter represeta a una entidad de sede en el sistema.
type Headquarter struct {
	Headquarter   uuid.UUID
	Name          string
	InstitutionID uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreatedBy     uuid.UUID
	UpdatedBy     uuid.UUID
}
