package customer

import (
	"context"
	"fmt"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/payload"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/domain"
)

// GetCustomers recupera una lista de clientes basándose en los criterios de búsqueda proporcionados.
func (s *Service) GetCustomers(ctx context.Context, criteria *payload.SearchCustomerCriteria) ([]domain.Customer, error) {
	customers, err := s.repository.GetCustomers(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("failed to get customers: %w", err)
	}

	if len(customers) == 0 {
		return customers, nil
	}

	return customers, nil
}
