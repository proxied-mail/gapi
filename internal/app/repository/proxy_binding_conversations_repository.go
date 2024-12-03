package repository

import (
	"errors"
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"time"
)

type ProxyBindingBotConversationsRepositoryInterface interface {
	GetLatest(pbBotId int, sender string) (models.ProxyBindingBotConversations, error)
	UpdateLastMessageReceived(pbBotId int, sender string) error
	UpdateLastMessageSent(pbBotId int, sender string) error
	CreateConversation(pbBotId int, sender string) (models.ProxyBindingBotConversations, error)
}

type ProxyBindingBotConversationsRepository struct {
	fx.In
	Db *gorm.DB
	ProxyBindingBotsRepositoryInterface
}

func (r ProxyBindingBotConversationsRepository) GetLatest(
	pbBotId int,
	sender string,
) (models.ProxyBindingBotConversations, error) {
	model := models.ProxyBindingBotConversations{}
	r.Db.Where("pb_bot_id", pbBotId).Where("sender_email", sender).Last(&model)
	if model.Id < 1 {
		return model, errors.New("cant find proxy binding conversation")
	}

	return model, nil
}

func (r ProxyBindingBotConversationsRepository) UpdateLastMessageReceived(
	pbBotId int,
	sender string,
) error {
	m, err := r.GetLatest(pbBotId, sender)
	if err != nil {
		return errors.New("cant find proxy binding conversation")
	}
	m.LastMessageAt = time.Now()
	m.ReceivedMessagesCount++
	r.Db.Save(&m)
	return nil
}

func (r ProxyBindingBotConversationsRepository) UpdateLastMessageSent(
	pbBotId int,
	sender string,
) error {
	m, err := r.GetLatest(pbBotId, sender)
	if err != nil {
		return errors.New("cant find proxy binding conversation")
	}
	m.LastMessageAt = time.Now()
	m.SentMessagesCount++
	r.Db.Save(&m)
	return nil
}

func (r ProxyBindingBotConversationsRepository) CreateConversation(
	pbBotId int,
	sender string,
) (models.ProxyBindingBotConversations, error) {
	model := models.ProxyBindingBotConversations{}

	proxyBindingBot, err := r.ProxyBindingBotsRepositoryInterface.GetById(pbBotId)
	if err != nil {
		return model, errors.New("cant find proxy binding bot")
	}

	model.ReceivedMessagesCount = 1
	model.SentMessagesCount = 0
	model.SenderEmail = sender
	model.ProxyBindingId = proxyBindingBot.ProxyBindingId
	model.PbBotId = proxyBindingBot.Id
	model.LastMessageAt = time.Now()
	r.Db.Save(&model)

	return model, nil
}
