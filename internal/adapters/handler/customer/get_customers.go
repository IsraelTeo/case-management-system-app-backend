package customer

import (
	"net/http"
	"strconv"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/payload"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/storage/specification"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/app/util"
	"github.com/labstack/echo/v4"
)

// GetCustomers maneja la solicitud HTTP para recuperar una lista de clientes basándose en los criterios de búsqueda proporcionados.
func (h *Handler) GetCustomers(c echo.Context) error {
	criteria := payload.SearchCustomerCriteria{
		BusinessOrName: util.BindOptionalString(c, "businessOrName"),
		TradeName:      util.BindOptionalString(c, "tradeName"),
		DNIOrRUC:       util.BindOptionalString(c, "dniOrRuc"),
		PhoneNumber:    util.BindOptionalString(c, "phoneNumber"),
		HomeAddress:    util.BindOptionalString(c, "homeAddress"),
		PersonalEmail:  util.BindOptionalString(c, "personalEmail"),
		Activity:       util.BindOptionalString(c, "activity"),
		Type:           util.BindOptionalString(c, "type"),
		BirthDate:      util.BindOptionalString(c, "birthDate"),
		CompanyID:      util.BindOptionalString(c, "companyId"),
		CreatedAt:      util.BindOptionalString(c, "createdAt"),
		SortField:      util.BindOptionalString(c, "sortField"),
		SortOrder:      util.BindOptionalString(c, "sortOrder"),
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err == nil {
		criteria.Page = page
	}

	size, err := strconv.Atoi(c.QueryParam("size"))
	if err == nil {
		criteria.Size = size
	}

	if err := specification.ValidateAndNormalizeSearchCriteria(&criteria); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	customers, err := h.service.GetCustomers(c.Request().Context(), &criteria)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, customers)
}
