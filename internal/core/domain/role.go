package domain

import (
	"time"

	"github.com/google/uuid"
)

// Role represents the "rol" table in the database.
type Role struct {
	ID              uuid.UUID
	Role            string
	Description     *string // TEXT, puede ser NULL → usamos puntero
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedByUserID uuid.UUID
	UpdatedByUserID uuid.UUID
}
