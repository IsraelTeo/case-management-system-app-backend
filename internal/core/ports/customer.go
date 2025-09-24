package ports

import (
	"context"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/payload"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CustomerRepository define los métodos para acceder y manipular clientes en la base de datos.
type CustomerRepository interface {
	Insert(ctx context.Context, c *domain.Customer) (*domain.Customer, error)
	GetByID(ctx context.Context, ID uuid.UUID) (*domain.Customer, error)
	Update(ctx context.Context, c *domain.Customer) (*domain.Customer, error)
	Delete(ctx context.Context, ID uuid.UUID) error
	GetCustomers(ctx context.Context, criteria *payload.SearchCustomerCriteria) ([]domain.Customer, error)
}

// CustomerService define los métodos de la lógica de negocio relacionados con clientes.
type CustomerService interface {
	Create(ctx context.Context, c *payload.CreateCustomerRequest, userID uuid.UUID) (*domain.Customer, error)
	GetByID(ctx context.Context, ID uuid.UUID) (*domain.Customer, error)
	Update(ctx context.Context, c *payload.UpdateCustomerRequest, customerID, userID uuid.UUID) (*domain.Customer, error)
	Delete(ctx context.Context, customerID uuid.UUID) error
	GetCustomers(ctx context.Context, criteria *payload.SearchCustomerCriteria) ([]domain.Customer, error)
}

// CustomerHandler define los métodos para manejar las solicitudes HTTP relacionadas con clientes.
type CustomerHandler interface {
	Create(c echo.Context) error
	GetByID(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetCustomers(c echo.Context) error
}
