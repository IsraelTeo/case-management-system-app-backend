package domain

import (
	"time"

	"github.com/google/uuid"
)

// User representa un usuario en el sistema.
type User struct {
	ID              uuid.UUID
	Username        string
	Password        string
	Email           string
	PhoneNumber     string
	RoleID          *uuid.UUID
	CompanyID       *uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedByUserID *uuid.UUID
	UpdatedByUserID *uuid.UUID
}
