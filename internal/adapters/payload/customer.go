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
	BirthDate      time.Time  `json:"birth_date" validate:"required"`
	CompanyID      *uuid.UUID `json:"company_id" validate:"omitempty,uuid"`
}
