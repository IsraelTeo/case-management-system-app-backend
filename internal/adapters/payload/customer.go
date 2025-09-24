package payload

import (
	"time"

	"github.com/google/uuid"
)

// CreateCustomerRequest representa el cuerpo JSON de la solicitud para crear un nuevo cliente.
type CreateCustomerRequest struct {
	BusinessOrName string     `json:"business_or_name" validate:"required,max=45"`
	TradeName      string     `json:"trade_name" validate:"required,max=45"`
	DNIOrRUC       string     `json:"dni_or_ruc" validate:"required,max=45"`
	PhoneNumber    string     `json:"phone_number" validate:"required,max=45"`
	HomeAddress    *string    `json:"home_address" validate:"omitempty,max=150"`
	PersonalEmail  string     `json:"personal_email" validate:"required,email,max=45"`
	Activity       string     `json:"activity" validate:"required,max=100"`
	Type           bool       `json:"type" validate:"required"`
	BirthDate      time.Time  `json:"birth_date" validate:"required,past"`
	CompanyID      *uuid.UUID `json:"company_id" validate:"omitempty,uuid"`
}

// UpdateCustomerRequest representa el cuerpo JSON de la solicitud para actualizar un cliente existente.
type UpdateCustomerRequest struct {
	BusinessOrName string     `json:"business_or_name" validate:"required,max=45"`
	TradeName      string     `json:"trade_name" validate:"required,max=45"`
	DNIOrRUC       string     `json:"dni_or_ruc" validate:"required,max=45"`
	PhoneNumber    string     `json:"phone_number" validate:"required,max=45"`
	HomeAddress    *string    `json:"home_address" validate:"omitempty,max=150"`
	PersonalEmail  string     `json:"personal_email" validate:"required,email,max=45"`
	Activity       string     `json:"activity" validate:"required,max=100"`
	Type           bool       `json:"type" validate:"required"`
	BirthDate      time.Time  `json:"birth_date" validate:"required"`
	CompanyID      *uuid.UUID `json:"company_id" validate:"omitempty,uuid"`
}

// SearchCustomerCriteria representa los criterios de búsqueda para filtrar clientes.
type SearchCustomerCriteria struct {
	BusinessOrName  *string `query:"businessOrName"`
	TradeName       *string `query:"tradeName"`
	DNIOrRUC        *string `query:"dniOrRuc"`
	PhoneNumber     *string `query:"phoneNumber"`
	HomeAddress     *string `query:"homeAddress"`
	PersonalEmail   *string `query:"personalEmail"`
	Activity        *string `query:"activity"`
	Type            *string `query:"type"`
	BirthDate       *string `query:"birthDate"`
	CreatedAt       *string `query:"createdAt"`
	CompanyID       *string `query:"companyId"`
	SortField       *string `query:"sortField"` // campo para ordenar
	SortOrder       *string `query:"sortOrder"` // dirección de ordenamiento "asc" o "desc"
	Page            int     `query:"page"`      // representa la página en la que se encuentra
	Size            int     `query:"size"`      // representa la cantidad de registros se mostrarán por página
	ParsedType      *bool
	ParsedBirthDate *time.Time
	ParsedCreatedAt *time.Time
	ParsedSortOrder *string
}
