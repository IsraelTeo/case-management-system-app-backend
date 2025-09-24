package customer

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

// Delete elimina un cliente de la base de datos por su ID.
func (r *Repository) Delete(ctx context.Context, ID uuid.UUID) error {
	query := r.db.QueryBuilder.Delete("cliente").
		Where(sq.Eq{"idcliente": ID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
