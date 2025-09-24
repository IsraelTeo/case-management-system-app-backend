package customer

import (
	"context"
	"fmt"
	"time"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/payload"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/domain"
	"github.com/google/uuid"
)

// Update actualiza la información de un cliente existente.
func (s *Service) Update(ctx context.Context, c *payload.UpdateCustomerRequest, customerID, userID uuid.UUID) (*domain.Customer, error) {
	customerExisting, err := s.repository.GetByID(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("customer not found: %w", err)
	}

	customerExisting.BusinessOrName = c.BusinessOrName
	customerExisting.TradeName = c.TradeName
	customerExisting.DNIOrRUC = c.DNIOrRUC
	customerExisting.PhoneNumber = c.PhoneNumber
	customerExisting.HomeAddress = c.HomeAddress
	customerExisting.PersonalEmail = c.PersonalEmail
	customerExisting.Activity = c.Activity
	customerExisting.Type = c.Type
	customerExisting.BirthDate = c.BirthDate
	customerExisting.CompanyID = c.CompanyID
	customerExisting.UpdatedAt = time.Now()
	customerExisting.UpdatedByUserID = userID

	customerUpdated, err := s.repository.Update(ctx, customerExisting)
	if err != nil {
		return nil, fmt.Errorf("failed to update customer: %w", err)
	}

	return customerUpdated, nil
}
