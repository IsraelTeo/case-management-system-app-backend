package routers

import (
	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/storage/postgres"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/app/config"
	"github.com/labstack/echo/v4"

	handler "github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/handler/auth"
	repository "github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/storage/postgres/repository/user"
	service "github.com/IsraelTeo/case-management-system-app-backend/internal/core/usecase/auth"
)

// setUpAuth configura las rutas relacionadas con la autenticación.
func setUpAuth(api *echo.Group, db *postgres.DB, cfg *config.JWT) {

	loginRepository := repository.New(db)
	loginService := service.New(loginRepository, cfg)
	loginHandler := handler.New(loginService)

	auth := api.Group("/auth")

	auth.POST(loginPath, loginHandler.Login)
}
