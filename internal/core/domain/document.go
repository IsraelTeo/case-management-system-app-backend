package domain

import (
	"time"

	"github.com/google/uuid"
)

// Document representa
type Document struct {
	ID                   uuid.UUID `json:"id"`
	DocumentNumber       string    `json:"document_number"`
	DigitalizationNumber string    `json:"digitalization_number"`
	NotificationNumber   string    `json:"notification_number,omitempty"`
	Type                 int       `json:"type"`
	Status               int       `json:"status"`
	CaseFileID           uuid.UUID `json:"case_file_id"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	CreatedByUserID      uuid.UUID `json:"created_by_user_id"`
	UpdatedByUserID      uuid.UUID `json:"updated_by_user_id"`
}
