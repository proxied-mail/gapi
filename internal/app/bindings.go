package app

import (
	"github.com/abrouter/gapi/internal/app/repository"
	"go.uber.org/fx"
)

func ProvideFxBindings() []fx.Option {
	return []fx.Option{
		fx.Provide(
			func(s repository.ProxyBindingBotMessagesRepository) repository.ProxyBindingBotMessagesRepositoryInterface {
				return s
			},
			func(s repository.ProxyBindingBotsRepository) repository.ProxyBindingBotsRepositoryInterface {
				return s
			},
			func(s repository.ReceivedEmailsRepository) repository.ReceivedEmailsRepositoryInterface {
				return s
			},
		),
	}
}
