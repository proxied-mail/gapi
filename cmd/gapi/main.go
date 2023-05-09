package main

import (
	"github.com/abrouter/gapi/internal/app"
	"github.com/abrouter/gapi/internal/app/provider"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			provider.EchoProvider,
		),
		fx.Invoke(
			app.ConfigureApiRoutes,
			app.StartHttpServer,
		),
	).Run()
}
