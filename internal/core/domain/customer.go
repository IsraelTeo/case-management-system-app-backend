package domain

import (
	"time"

	"github.com/google/uuid"
)

// Customer representa una entidad de cliente en el sistema.
type Customer struct {
	ID              uuid.UUID
	BusinessName    string
	TradeName       string
	DNIRUC          string
	HomeAddress     *string
	PersonalEmail   string
	Type            bool
	BirthDate       time.Time
	CompanyID       *uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedByUserID uuid.UUID
	UpdatedByUserID uuid.UUID
}
