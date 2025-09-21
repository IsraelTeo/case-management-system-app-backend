package config

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func StartServer(e *echo.Echo, port string) error {
	err := e.Start(":" + port)
	if err != nil {
		return fmt.Errorf("error starting server on port %s: %w", port, err)
	}

	return nil
}
