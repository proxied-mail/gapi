package app

import (
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/bot_messages"
	"github.com/abrouter/gapi/internal/app/services/conversations"
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
			func(
				s repository.ProxyBindingBotConversationsRepository,
			) repository.ProxyBindingBotConversationsRepositoryInterface {
				return s
			},
			func(s conversations.ConversationManager) conversations.ConversationManagerInterface {
				return s
			},
			func(s bot_messages.MessageSaverService) bot_messages.MessageSaverServiceInterface {
				return s
			},
			func(s repository.BotsRepository) repository.BotsRepositoryInterface {
				return s
			},
		),
	}
}
