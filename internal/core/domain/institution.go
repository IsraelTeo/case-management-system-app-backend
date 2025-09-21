package domain

import (
	"time"

	"github.com/google/uuid"
)

// Institution representa a una entidad de institución en el sistema.
type Institution struct {
	InstitutionID uuid.UUID
	Name          string
	Description   string
	UpdatedAt     time.Time
	CreatedBy     uuid.UUID
	UpdatedBy     uuid.UUID
}
