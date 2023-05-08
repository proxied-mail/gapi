package main

import (
	"github.com/abrouter/gapi/internal/app/provider"
	"go.uber.org/fx"
)
import "github.com/abrouter/gapi/internal/app"

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
