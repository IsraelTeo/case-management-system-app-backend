package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/IsraelTeo/case-management-system-app-backend/internal/adapters/storage/postgres"
	"github.com/IsraelTeo/case-management-system-app-backend/internal/app/config"
)

func main() {
	// Cargar Variables de Entorno
	cfg, err := config.New()
	if err != nil {
		slog.Error("Error loading configuration", "error", err)
		return
	}

	ctx := context.Background()

	// Conectar a la base de datos
	db, err := postgres.New(ctx, cfg.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfully connected to the database", "db", cfg.DB.Name)

	// Migrar tablas
	err = db.Migrate(cfg.DB)
	if err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfully migrated the database")
}
