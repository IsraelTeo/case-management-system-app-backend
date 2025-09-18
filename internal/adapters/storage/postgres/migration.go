package postgres

import (
	"embed"
	"fmt"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/app/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

/*
* Variable que guarda es un "contenedor" dentro del binario
* que guarda todos los archivos .sql de la carpeta migrations/.
* Así la aplicación puede ejecutar migraciones sin depender de
* archivos externos en el sistema.
 */
//go:embed migrations/*.sql
var migrationsFS embed.FS

// Migrate ejecuta las migraciones de la base de datos
func (db *DB) Migrate(cfg *config.DB) error {
	driver, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return err
	}

	// construye la url directamente desde cfg
	url := buildMigrateURL(cfg)

	migrations, err := migrate.NewWithSourceInstance("iofs", driver, url)
	if err != nil {
		return err
	}

	if err := migrations.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func buildMigrateURL(cfg *config.DB) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
	)
}
