package repositories

import (
	"go.uber.org/fx"
)

// Module provides the repositories module for FX
func Module() fx.Option {
	return fx.Module("repositories",
		fx.Provide(NewUserRepository),
	)
}
