package repository

import (
	"database/sql"
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type ProxyBindingBotMessagesRepositoryInterface interface {
	Create(
		pbBot models.ProxyBindingBots,
		receivedEmail int,
		sender string,
		conversationId int,
	) error
}

type ProxyBindingBotMessagesRepository struct {
	fx.In
	Db *gorm.DB
}

func (c ProxyBindingBotMessagesRepository) Create(
	pbBot models.ProxyBindingBots,
	receivedEmail int,
	sender string,
	conversationId int,
) error {
	model := models.ProxyBindingBotMessages{
		ReceivedEmailId: receivedEmail,
		PbBotId:         pbBot.Id,
		ProxyBindingId:  pbBot.ProxyBindingId,
		SenderEmail:     sender,
		ConversationId: sql.NullInt64{
			Int64: int64(conversationId),
		},
	}
	c.Db.Save(&model)

	return nil
}
