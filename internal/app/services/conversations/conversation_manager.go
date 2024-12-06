package conversations

import (
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/internal/app/repository"
	"go.uber.org/fx"
	"time"
)

type ConversationManagerInterface interface {
	MessageReceived(
		pbBot models.ProxyBindingBots,
		email models.ReceivedEmails,
	) (models.ProxyBindingBotConversations, error)
}

type ConversationManager struct {
	fx.In
	repository.ProxyBindingBotConversationsRepositoryInterface
}

func (cm ConversationManager) MessageReceived(
	pbBot models.ProxyBindingBots,
	email models.ReceivedEmails,
) (models.ProxyBindingBotConversations, error) {
	lastConv, _ := cm.ProxyBindingBotConversationsRepositoryInterface.GetLatest(
		pbBot.Id,
		email.SenderEmail,
	)
	if lastConv.Id < 1 {
		conv, err := cm.CreateConversation(pbBot.Id, email.SenderEmail)
		if err != nil {
			return models.ProxyBindingBotConversations{}, err
		}

		return conv, nil
	}
	lastMessageDate := lastConv.LastMessageAt
	isExpired := time.Now().Unix() > lastMessageDate.Unix()+(int64(pbBot.SessionLength)*60)

	if isExpired {
		cm.ProxyBindingBotConversationsRepositoryInterface.DeactivateConversations(pbBot.Id, email.SenderEmail)
		conv, err := cm.CreateConversation(pbBot.Id, email.SenderEmail)
		if err != nil {
			return models.ProxyBindingBotConversations{}, err
		}
		return conv, nil
	}

	err3 := cm.ProxyBindingBotConversationsRepositoryInterface.UpdateLastMessageReceived(
		pbBot.Id,
		email.SenderEmail,
	)

	_ = cm.ProxyBindingBotConversationsRepositoryInterface.DeactivateConversationsExcept(
		pbBot.Id,
		email.SenderEmail,
		lastConv.Id,
	)

	if err3 != nil {
		return lastConv, err3
	}

	return lastConv, nil
}
