package database

import (
	"context"

	"github.com/rs/zerolog"
	"go.uber.org/fx"

	"github.com/dimashky/go-appointment/app/config"
	"github.com/dimashky/go-appointment/app/models"
)

// DatabaseConfigProvider extracts database config from main config
func DatabaseConfigProvider(cfg config.Config) *config.Database {
	return &cfg.Database
}

// Module provides the database module for FX
func Module() fx.Option {
	return fx.Module("database",
		fx.Provide(DatabaseConfigProvider),
		fx.Provide(NewDB),
		fx.Invoke(registerHooks),
	)
}

// registerHooks sets up lifecycle hooks for the database
func registerHooks(lc fx.Lifecycle, db *DB, logger zerolog.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info().Msg("Database connection established")
			// Run auto-migration
			logger.Info().Msg("Running database migrations...")
			if err := models.AutoMigrate(db.DB); err != nil {
				logger.Error().Err(err).Msg("Failed to run migrations")
				return err
			}
			logger.Info().Msg("Database migrations completed successfully")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info().Msg("Closing database connection")
			return db.Close()
		},
	})
}
