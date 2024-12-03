package repository

import (
	"errors"
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type ProxyBindingBotMessagesRepositoryInterface interface {
	Create(
		receivedEmail int,
		pbBotId int,
	) error
}

type ProxyBindingBotMessagesRepository struct {
	fx.In
	ProxyBindingBotsRepositoryInterface
	ReceivedEmailsRepositoryInterface
	Db *gorm.DB
}

func (c ProxyBindingBotMessagesRepository) Create(receivedEmail int, pbBotId int) error {
	pbBot, pbBotErr := c.ProxyBindingBotsRepositoryInterface.getById(pbBotId)
	if pbBotErr != nil {
		return errors.New("cant find bot")
	}

	receivedEmailModel, errorReceived := c.ReceivedEmailsRepositoryInterface.getOneById(receivedEmail)
	if errorReceived != nil {
		return errors.New("Cannot find the received email")
	}

	model := models.ProxyBindingBotMessages{
		ReceivedEmailId: receivedEmail,
		PbBotId:         pbBotId,
		ProxyBindingId:  pbBot.ProxyBindingId,
		SenderEmail:     receivedEmailModel.SenderEmail,
	}
	c.Db.Save(&model)

	return nil
}
