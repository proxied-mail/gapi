package conversation_messages_rsp

import (
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/internal/app/services/received_emalis"
	"github.com/abrouter/gapi/pkg/entityId"
	"go.uber.org/fx"
)

type ConversationMessagesResponse struct {
	Messages []ConversationMessagesResponseMessage `json:"messages"`
	LastId   string                                `json:"lastId"`
}

type ConversationMessagesResponseMessage struct {
	Id             string                           `json:"id"`
	ConversationId string                           `json:"conversationId"`
	Message        received_emalis.ReceivedEmailDTO `json:"message"`
}

type ConversationMessagesTransformer struct {
	fx.In
	entityId.Encoder
	received_emalis.ReceivedEmailParser
}

func (cmt ConversationMessagesTransformer) Transform(
	models []models.ProxyBindingBotMessages,
	receivedEmails map[int]models.ReceivedEmails,
) ConversationMessagesResponse {
	response := ConversationMessagesResponse{}
	for _, model := range models {
		id := cmt.Encoder.Encode(model.Id, "proxy_binding_bot_messages")
		conversationId := cmt.Encoder.Encode(
			int(model.ConversationId.Int64),
			"proxy_binding_bot_conversations",
		)
		receivedEmail := receivedEmails[model.ReceivedEmailId]
		transformedReceivedEmail, err := cmt.ReceivedEmailParser.Parse(receivedEmail)
		if err != nil {
			continue
		}

		response.Messages = append(response.Messages, ConversationMessagesResponseMessage{
			Id:             id,
			ConversationId: conversationId,
			Message:        transformedReceivedEmail,
		})
	}

	return response
}
