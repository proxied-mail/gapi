package main

import (
	"github.com/abrouter/gapi/internal/app"
	"github.com/abrouter/gapi/internal/app/boot"
	"github.com/abrouter/gapi/internal/app/env"
	"github.com/abrouter/gapi/internal/app/provider"
	"github.com/abrouter/gapi/pkg/entityId"
	"go.uber.org/fx"
)

func main() {

	MysqlRwConnectionProvider := provider.MysqlRwConnectionProvider{}

	fx.New(
		fx.Provide(
			provider.EchoProvider,
			MysqlRwConnectionProvider.Connect,
			provider.OrmProvider,
			func() entityId.Encoder {
				return entityId.Encoder{}
			},
		),
		fx.Invoke(
			boot.ParseFlags,
			env.ReadEnv,
			MysqlRwConnectionProvider.Connect,
			app.ConfigureApiRoutes,
			app.StartHttpServer,
		),
	).Run()
}
