package domain

import (
	"time"

	"github.com/google/uuid"
)

// Role representa una entidad de rol en el sistema.
type Role struct {
	ID              uuid.UUID
	Role            string
	Description     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedByUserID uuid.UUID
	UpdatedByUserID uuid.UUID
}
