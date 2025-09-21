package routers

import (
	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/storage/postgres"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/app/config"
	"github.com/labstack/echo/v4"
)

// Definición de rutas comunes
const (
	idPath    = "/:id"
	voidPath  = ""
	dniPath   = "/dni"
	loginPath = "/login"
)

// InitEnpoints inicializa las rutas del API.
func InitEnpoints(e *echo.Echo, db *postgres.DB, cfg *config.JWT) {
	api := e.Group("/api/v1")
	setUpCustomer(api, db, cfg)
	setUpAuth(api, db, cfg)
}
