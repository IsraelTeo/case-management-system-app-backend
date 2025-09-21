package customer

import (
	"context"
	"fmt"
	"time"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/payload"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/domain"
	"github.com/google/uuid"
)

// Create crea un nuevo cliente.
func (s *Service) Create(ctx context.Context, c *payload.CreateCustomerRequest, userID uuid.UUID) (*domain.Customer, error) {
	ID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate customer ID: %w", err)
	}

	customer := &domain.Customer{
		ID:              ID,
		BusinessOrName:  c.BusinessOrName,
		TradeName:       c.TradeName,
		DNIOrRUC:        c.DNIOrRUC,
		PhoneNumber:     c.PhoneNumber,
		HomeAddress:     c.HomeAddress,
		PersonalEmail:   c.PersonalEmail,
		Activity:        c.Activity,
		Type:            c.Type,
		BirthDate:       c.BirthDate,
		CompanyID:       c.CompanyID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		CreatedByUserID: userID,
		UpdatedByUserID: userID,
	}

	customerSaved, err := s.repository.Insert(ctx, customer)
	if err != nil {
		return nil, fmt.Errorf("failed to create customer: %w", err)
	}

	return customerSaved, nil
}
