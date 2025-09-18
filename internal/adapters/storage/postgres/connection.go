package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/app/config"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

/*
* DB es un contenedor para la conexión a bases de datos PostgreSQL.
* que utiliza pgxpool como controlador de base de datos.
* También contiene una referencia a squirrel.StatementBuilderType.
* que se utiliza para crear consultas SQL compatibles con la sintaxis
* PostgreSQL.
 */
type DB struct {
	pool         *pgxpool.Pool
	QueryBuilder *squirrel.StatementBuilderType
	url          string
}

// New crea una nueva instancia de DB
func New(ctx context.Context, cfg *config.DB) (*DB, error) {
	url := buildURL(cfg)

	log.Println("DB URL:", url)

	db, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &DB{
		db,
		&psql,
		url,
	}, nil
}

// buildURL construye la URL de conexión PostgreSQL con pool configurado
func buildURL(cfg *config.DB) string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s pool_min_conns=%d pool_max_conns=%d",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
		cfg.MinConn,
		cfg.MaxConn,
	)
}
