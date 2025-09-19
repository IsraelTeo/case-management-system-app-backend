package domain

import (
	"time"

	"github.com/google/uuid"
)

// Company representa una entidad de empresa en el sistema.
type Company struct {
	ID              uuid.UUID
	Name            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedByUserID uuid.UUID
	UpdatedByUserID uuid.UUID
}
