package domain

import (
	"time"

	"github.com/google/uuid"
)

// el booleano representa dos tipos de cliente, preguntar a madre:

// Customer representa una entidad de cliente en el sistema.
type Customer struct {
	ID              uuid.UUID
	BusinessOrName  string
	TradeName       string
	DNIOrRUC        string
	PhoneNumber     string
	HomeAddress     *string
	PersonalEmail   string
	Activity        string
	Type            bool
	BirthDate       time.Time
	CompanyID       *uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedByUserID uuid.UUID
	UpdatedByUserID uuid.UUID
}
