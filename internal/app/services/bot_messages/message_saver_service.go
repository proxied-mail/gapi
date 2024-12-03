package bot_messages

import (
	"errors"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/conversations"
	"go.uber.org/fx"
)

type MessageSaverServiceDto struct {
	ReceivedEmailId   int
	ProxyBindingBotId int
}

type MessageSaverServiceInterface interface {
	Save(m MessageSaverServiceDto) error
}

type MessageSaverService struct {
	fx.In
	repository.ProxyBindingBotConversationsRepositoryInterface
	repository.ProxyBindingBotMessagesRepositoryInterface
	repository.ReceivedEmailsRepositoryInterface
	conversations.ConversationManagerInterface
	repository.ProxyBindingBotsRepositoryInterface
}

func (m MessageSaverService) Save(dto MessageSaverServiceDto) error {
	receivedEmail, err := m.ReceivedEmailsRepositoryInterface.GetOneById(dto.ReceivedEmailId)
	if err != nil {
		return errors.New("cant find received email")
	}

	pbBot, err2 := m.ProxyBindingBotsRepositoryInterface.GetById(dto.ProxyBindingBotId)
	if err2 != nil {
		return errors.New("Cant find proxy binding bot")
	}

	conv, err3 := m.ConversationManagerInterface.MessageReceived(pbBot, receivedEmail)

	if err3 != nil {
		return errors.New("Failed to create the conversation")
	}

	m.ProxyBindingBotMessagesRepositoryInterface.Create(
		pbBot,
		receivedEmail.Id,
		receivedEmail.SenderEmail,
		conv.Id,
	)

	return nil
}