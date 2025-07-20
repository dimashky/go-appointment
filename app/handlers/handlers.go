package handlers

import (
	"github.com/CodeLieutenant/uberfx-common/v3/http/fiber/fiberfx"

	"github.com/dimashky/go-appointment/app/handlers/helloworld"
)

func Handlers() fiberfx.RoutesFx {
	return fiberfx.Routes(
		[]fiberfx.RouteFx{
			fiberfx.Get("/", helloworld.HelloWorld),
		},
		fiberfx.WithPrefix("/"),
	)
}
