package customer

import (
	"context"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/domain"
)

// Insert agrega un nuevo cliente a la tabla "cliente"
func (r *Repository) Insert(ctx context.Context, c *domain.Customer) (*domain.Customer, error) {
	// Insert(): Construye la consulta insert con los campos del cliente
	query := r.db.QueryBuilder.Insert("cliente").
		Columns(
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
		Values(
			c.ID,
			c.BusinessOrName,
			c.TradeName,
			c.DNIOrRUC,
			c.PhoneNumber,
			c.HomeAddress,
			c.PersonalEmail,
			c.Activity,
			c.Type,
			c.BirthDate,
			c.CompanyID,
			c.CreatedAt,
			c.UpdatedAt,
			c.CreatedByUserID,
			c.UpdatedByUserID,
		).
		Suffix("RETURNING *")

	// ToSql(): Convierte la consulta en una cadena SQL lista para ejecutar
	// y un slice con los argumentos (args) que reemplazarán los placeholders.
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	// QueryRow():Ejecuta la consulta SQL
	// Scan(): Escanea la fila devuelta y actualiza los campos del struct Customer con los valores que tiene la base de datos.
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
		&c.CreatedAt,
		&c.UpdatedAt,
		&c.CreatedByUserID,
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
