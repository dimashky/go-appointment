package handlers

import (
	"github.com/CodeLieutenant/uberfx-common/v3/http/fiber/fiberfx"

	"github.com/dimashky/go-appointment/app/handlers/users"
)

func Handlers() fiberfx.RoutesFx {
	return fiberfx.Routes(
		[]fiberfx.RouteFx{
			fiberfx.Post("/api/users", users.CreateUserRoute),
			fiberfx.Get("/api/users/:id", users.GetUserRoute),
			fiberfx.Get("/api/users", users.ListUsersRoute),
		},
		fiberfx.WithPrefix("/"),
	)
}
