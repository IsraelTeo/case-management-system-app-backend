package domain

import (
	"time"

	"github.com/google/uuid"
)

// Lawyer representa una entidad de abogado en el sistema.
type Lawyer struct {
	ID              uuid.UUID
	FirstName       string
	LastName        string
	DNI             string
	LicenseNumber   string
	CellPhone       string
	OfficePhone     string
	HomePhone       string
	HomeAddress     string
	PersonalEmail   string
	OfficeEmail     string
	BirthDate       time.Time
	CompanyID       uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedByUserID uuid.UUID
	UpdatedByUserID uuid.UUID
}
