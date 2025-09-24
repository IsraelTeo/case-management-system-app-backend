package customer

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// Delete elimina un cliente por su ID.
func (s *Service) Delete(ctx context.Context, customerID uuid.UUID) error {
	_, err := s.repository.GetByID(ctx, customerID)
	if err != nil {
		return fmt.Errorf("customer not found: %w", err)
	}

	if err := s.repository.Delete(ctx, customerID); err != nil {
		return fmt.Errorf("failed to delete customer: %w", err)
	}

	return nil
}
