package routers

import (
	handler "github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/handler/customer"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/storage/postgres"
	repository "github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/storage/postgres/repository/customer"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/app/config"
	service "github.com/IsraelTeo/case-management-system-app-backend/internal/core/usecase/customer"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/middleware"
	"github.com/labstack/echo/v4"
)

// setUpCustomer configura las rutas relacionadas con clientes.
func setUpCustomer(api *echo.Group, db *postgres.DB, cfg *config.JWT) {
	customerRepository := repository.New(db)
	customerService := service.New(customerRepository)
	customerHandler := handler.New(customerService, cfg)

	customer := api.Group("/customers")

	customer.POST(voidPath, middleware.ValidateJWT(customerHandler.Create, cfg))
}
