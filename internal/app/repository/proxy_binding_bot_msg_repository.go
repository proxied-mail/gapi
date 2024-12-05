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
	Query(pbBotId int, lastProxyBinding int) []models.ProxyBindingBotMessages
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
			Valid: true,
		},
	}
	c.Db.Save(&model)

	return nil
}

func (c ProxyBindingBotMessagesRepository) Query(
	pbBotId int,
	lastProxyBinding int,
) []models.ProxyBindingBotMessages {
	q := c.Db.Where("proxy_binding_bot_messages.pb_bot_id", pbBotId)
	if lastProxyBinding > 0 {
		q = q.Where("proxy_binding_bot_messages.id > ?", lastProxyBinding)
	}
	q = q.Joins(
		"JOIN proxy_binding_bot_conversations pbbc on proxy_binding_bot_messages.conversation_id = pbbc.id ",
	)
	q.Where("pbbc.status = ?", 1)

	var modelsList []models.ProxyBindingBotMessages
	q.Find(&modelsList)

	return modelsList
}
