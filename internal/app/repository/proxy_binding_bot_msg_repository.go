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
	Query(pbBotId int, lastProxyBinding int, onlyUnread bool) []models.ProxyBindingBotMessages
	QueryByBotUid(pbBotUid string, lastProxyBinding int, onlyUnread bool) []models.ProxyBindingBotMessages
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
	onlyUnread bool,
) []models.ProxyBindingBotMessages {
	q := c.Db.Where("proxy_binding_bot_messages.pb_bot_id", pbBotId)
	if lastProxyBinding > 0 {
		q = q.Where("proxy_binding_bot_messages.id > ?", lastProxyBinding)
	}
	q = q.Joins(
		"JOIN proxy_binding_bot_conversations pbbc on proxy_binding_bot_messages.conversation_id = pbbc.id ",
	)
	q = q.Joins(
		"Join proxy_binding_bots pbp on pbp.bot_id=pbbc.pb_bot_id",
	)
	q = q.Where("pbp.status = ?", models.PB_BOT_STATUS_ACTIVE)

	if onlyUnread {
		q = q.Where("proxy_binding_bot_messages.read = ?", false)
	}

	var modelsList []models.ProxyBindingBotMessages
	q.Find(&modelsList)

	return modelsList
}

func (c ProxyBindingBotMessagesRepository) QueryByBotUid(
	pbBotId string,
	lastProxyBinding int,
	onlyUnread bool,
) []models.ProxyBindingBotMessages {
	q := c.Db.Table("proxy_binding_bot_messages").Where("bots.uid = ?", pbBotId)
	q = q.Joins(
		"JOIN proxy_binding_bot_conversations pbbc on proxy_binding_bot_messages.conversation_id = pbbc.id",
	)
	q = q.Joins(
		"Join proxy_binding_bots pbp on pbp.bot_id=pbbc.pb_bot_id",
	)
	q = q.Joins("join bots on bots.id = pbp.bot_id")
	q = q.Where("pbp.status = ?", models.PB_BOT_STATUS_ACTIVE)

	q.Where("pbbc.status = ?", 1)
	if lastProxyBinding > 0 {
		q = q.Where("proxy_binding_bot_messages.id > ?", lastProxyBinding)
	}
	if onlyUnread {
		q = q.Where("proxy_binding_bot_messages.read = ?", false)
	}

	var modelsList []models.ProxyBindingBotMessages
	q.Find(&modelsList)

	return modelsList
}
