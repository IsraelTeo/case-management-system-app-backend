package customer

import (
	"context"
	"fmt"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/domain"
	"github.com/google/uuid"
)

// GetByID obtiene un cliente por su ID.
func (s *Service) GetByID(ctx context.Context, ID uuid.UUID) (*domain.Customer, error) {
	var customer *domain.Customer

	customer, err := s.repository.GetByID(ctx, ID)
	if err != nil {
		return nil, fmt.Errorf("customer not found: %w", err)
	}

	return customer, nil
}
