package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/core/domain"
)

// GetByEmail obtiene un usuario por su email desde la base de datos.
func (r *Repository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User

	query := r.db.QueryBuilder.Select(
		"idusuario",
		"usuario",
		"password",
		"email",
		"telefono",
		"idrol",
		"idempresa",
		"fecha_registro",
		"fecha_actualizacion",
		"usuario_registro",
		"usuario_actualizacion",
	).
		From("public.usuario").
		Where(sq.Eq{"usuario": username}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.PhoneNumber,
		&user.RoleID,
		&user.CompanyID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.CreatedByUserID,
		&user.UpdatedByUserID,
	)

	fmt.Printf("DEBUG user después de Scan: %+v\n", user)

	if err != nil {
		fmt.Printf("ERROR en Scan: %v\n", err)
		if err == pgx.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}
