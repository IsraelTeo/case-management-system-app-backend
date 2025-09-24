package customer

import (
	"context"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/domain"
)

func (r *Repository) Update(ctx context.Context, c *domain.Customer) (*domain.Customer, error) {
	query := r.db.QueryBuilder.Update("cliente").
		Set("nombre_razon", c.BusinessOrName).
		Set("apellido_comercial", c.TradeName).
		Set("dni_ruc", c.DNIOrRUC).
		Set("telefono_celular", c.PhoneNumber).
		Set("direccion_domicilio", c.HomeAddress).
		Set("email_personal", c.PersonalEmail).
		Set("actividad", c.Activity).
		Set("tipo", c.Type).
		Set("fecha_nacimiento", c.BirthDate).
		Set("idempresa", c.CompanyID).
		Set("fecha_actualizacion", c.UpdatedAt).
		Set("usuario_actualizacion", c.UpdatedByUserID).
		Where("idcliente = ?", c.ID).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&c.ID,
		&c.BusinessOrName,
		&c.TradeName,
		&c.DNIOrRUC,
		&c.PhoneNumber,
		&c.HomeAddress,
		&c.PersonalEmail,
		&c.Activity,
		&c.Type,
		&c.BirthDate,
		&c.CompanyID,
		&c.UpdatedAt,
		&c.UpdatedByUserID,
	)
	if err != nil {
		if errCode := r.db.ErrorCode(err); errCode == "23505" {
			return nil, err
		}

		return nil, err
	}

	return c, nil
}
