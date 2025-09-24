package util

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// bindOptionalString es una función auxiliar para enlazar parámetros de consulta opcionales de tipo string.
func BindOptionalString(c echo.Context, param string) *string {
	if v := c.QueryParam(param); v != "" {
		return &v
	}

	return nil
}

// BindOptionalInt es una función auxiliar para enlazar parámetros de consulta opcionales de tipo int.
func BindOptionalInt(c echo.Context, param string) *int {
	if v := c.QueryParam(param); v != "" {
		if intValue, err := strconv.Atoi(v); err == nil {
			return &intValue
		}
	}

	return nil
}
