package specification

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/payload"
)

// ValidateAndNormalizeSearchCriteria valida y convierte los parámetros de búsqueda
func ValidateAndNormalizeSearchCriteria(criteria *payload.SearchCustomerCriteria) error {
	parsedBirthDate, err := validateDate(criteria.BirthDate)
	if err != nil {
		return err
	}
	criteria.ParsedBirthDate = parsedBirthDate

	parsedCreatedAt, err := validateDate(criteria.CreatedAt)
	if err != nil {
		return err
	}
	criteria.ParsedCreatedAt = parsedCreatedAt

	if err := validateType(criteria); err != nil {
		return err
	}

	if err := validateSort(criteria); err != nil {
		return err
	}

	page, size := validatePagination(criteria.Page, criteria.Size)
	criteria.Page = page
	criteria.Size = size

	return nil
}

// validateDate valida y convierte una cadena de fecha en un puntero de time.Time
func validateDate(dateStr *string) (*time.Time, error) {
	if dateStr == nil || *dateStr == "" {
		return nil, nil
	}

	t, err := time.Parse("2006-01-02", *dateStr)
	if err != nil {
		return nil, errors.New("invalid date format, expected YYYY-MM-DD")
	}
	return &t, nil
}

// validateType valida y normaliza el campo Type
func validateType(criteria *payload.SearchCustomerCriteria) error {
	if criteria.Type != nil {
		b, err := strconv.ParseBool(*criteria.Type)
		if err != nil {
			return errors.New("invalid type value, expected true or false")
		}
		criteria.ParsedType = &b
	}
	return nil
}

// validatePagination valida y normaliza los parámetros de paginación
func validatePagination(page, size int) (int, int) {
	if size <= 0 {
		size = 10
	}

	if page < 0 {
		page = 0
	}

	return page, size
}

// validateSort valida los parámetros de ordenamiento
func validateSort(criteria *payload.SearchCustomerCriteria) error {
	if criteria.SortField != nil && criteria.SortOrder != nil {
		order := strings.ToLower(*criteria.SortOrder)
		if order != "asc" && order != "desc" {
			return errors.New("sortOrder must be 'asc' or 'desc'")
		}

		criteria.ParsedSortOrder = &order
	}

	return nil
}
