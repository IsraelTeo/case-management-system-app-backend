package customer

import (
	"context"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/domain"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

// GetByID obtiene un cliente por su ID desde la base de datos.
func (r *Repository) GetByID(ctx context.Context, ID uuid.UUID) (*domain.Customer, error) {
	var customer domain.Customer

	query := r.db.QueryBuilder.Select(
		"idcliente",
		"nombre_razon",
		"apellido_comercial",
		"dni_ruc",
		"telefono_celular",
		"direccion_domicilio",
		"email_personal",
		"actividad",
		"tipo",
		"fecha_nacimiento",
		"idempresa",
		"fecha_registro",
		"fecha_actualizacion",
		"usuario_registro",
		"usuario_actualizacion",
	).
		From("cliente").
		Where(sq.Eq{"idcliente": ID}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&customer.ID,
		&customer.BusinessOrName,
		&customer.TradeName,
		&customer.DNIOrRUC,
		&customer.PhoneNumber,
		&customer.HomeAddress,
		&customer.PersonalEmail,
		&customer.Activity,
		&customer.Type,
		&customer.BirthDate,
		&customer.CompanyID,
		&customer.CreatedAt,
		&customer.UpdatedAt,
		&customer.CreatedByUserID,
		&customer.UpdatedByUserID,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	return &customer, nil
}
