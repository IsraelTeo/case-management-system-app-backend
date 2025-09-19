package domain

import (
	"time"

	"github.com/google/uuid"
)

// CaseFile representa a una entidad de expediente judicial en el sistema.
type CaseFile struct {
	CaseFileID uuid.UUID
	Status     int
	FileNumber string
	Subject    string
	BranchID   uuid.UUID
	CompanyID  uuid.UUID
	ClientID   uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CreatedBy  uuid.UUID
	UpdatedBy  uuid.UUID
}
