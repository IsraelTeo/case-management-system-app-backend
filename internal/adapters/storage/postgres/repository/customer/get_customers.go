package customer

import (
	"context"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/payload"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/domain"
	sq "github.com/Masterminds/squirrel"
)

// GetCustomers recupera una lista de clientes basándose en los criterios de búsqueda proporcionados.
func (r *Repository) GetCustomers(ctx context.Context, criteria *payload.SearchCustomerCriteria) ([]domain.Customer, error) {
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
	).From("cliente")

	if criteria.BusinessOrName != nil {
		query = query.Where(sq.ILike{"nombre_razon": "%" + *criteria.BusinessOrName + "%"})
	}

	if criteria.TradeName != nil {
		query = query.Where(sq.ILike{"apellido_comercial": "%" + *criteria.TradeName + "%"})
	}

	if criteria.DNIOrRUC != nil {
		query = query.Where(sq.ILike{"dni_ruc": "%" + *criteria.DNIOrRUC + "%"})
	}

	if criteria.PhoneNumber != nil {
		query = query.Where(sq.Eq{"telefono_celular": *criteria.PhoneNumber})
	}

	if criteria.HomeAddress != nil {
		query = query.Where(sq.ILike{"direccion_domicilio": "%" + *criteria.HomeAddress + "%"})
	}

	if criteria.PersonalEmail != nil {
		query = query.Where(sq.Eq{"email_personal": *criteria.PersonalEmail})
	}

	if criteria.Activity != nil {
		query = query.Where(sq.ILike{"actividad": "%" + *criteria.Activity + "%"})
	}

	if criteria.ParsedBirthDate != nil {
		query = query.Where(sq.Eq{"fecha_nacimiento": *criteria.ParsedBirthDate})
	}

	if criteria.ParsedCreatedAt != nil {
		query = query.Where(sq.Eq{"fecha_registro": *criteria.ParsedCreatedAt})
	}

	if criteria.ParsedType != nil {
		query = query.Where(sq.Eq{"tipo": *criteria.ParsedType})
	}

	if criteria.ParsedSortOrder != nil {
		query = query.OrderBy(*criteria.SortField + " " + *criteria.ParsedSortOrder)
	}

	query = applyPagination(query, criteria)

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Pool.Query(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []domain.Customer
	for rows.Next() {
		var c domain.Customer
		err := rows.Scan(
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
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}

// applyPagination aplica limit y offset al query según los criterios de búsqueda.
func applyPagination(query sq.SelectBuilder, criteria *payload.SearchCustomerCriteria) sq.SelectBuilder {
	offset := criteria.Page * criteria.Size
	return query.Limit(uint64(criteria.Size)).Offset(uint64(offset))
}
