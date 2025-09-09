package users

import (
	"go.uber.org/fx"
)

// Module provides the user handlers module for FX
func Module() fx.Option {
	return fx.Module("user-handlers",
		fx.Provide(NewHandler),
	)
}
