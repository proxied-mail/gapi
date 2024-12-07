package conversation_messages_rsp

import (
	"encoding/json"
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/internal/app/services/received_emalis"
	"github.com/abrouter/gapi/pkg/entityId"
	"go.uber.org/fx"
)

type ConversationMessagesResponse struct {
	ProxyBindingBots map[string]ProxyBindingBotResponse    `json:"proxyBindingBots"`
	Messages         []ConversationMessagesResponseMessage `json:"messages"`
	LastId           string                                `json:"lastId"`
}

type ProxyBindingBotResponse struct {
	Id            string      `json:"id"`
	Status        int         `json:"status"`
	SessionLength int         `json:"session_length"`
	Config        interface{} `json:"config"`
}

type ConversationMessagesResponseMessage struct {
	Id                string                           `json:"id"`
	ConversationId    string                           `json:"conversationId"`
	Read              bool                             `json:"read"`
	Message           received_emalis.ReceivedEmailDTO `json:"message"`
	ProxyBindingId    string                           `json:"proxyBindingId"`
	ProxyBindingBotId string                           `json:"proxyBindingBotId"`
}

type ConversationMessagesTransformer struct {
	fx.In
	entityId.Encoder
	received_emalis.ReceivedEmailParser
}

func (cmt ConversationMessagesTransformer) Transform(
	models []models.ProxyBindingBotMessages,
	receivedEmails map[int]models.ReceivedEmails,
	pbBots []models.ProxyBindingBots,
) ConversationMessagesResponse {
	response := ConversationMessagesResponse{}

	var id string
	for _, model := range models {
		id = cmt.Encoder.Encode(model.Id, "proxy_binding_bot_messages")
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
			Id:                id,
			ConversationId:    conversationId,
			Message:           transformedReceivedEmail,
			Read:              model.Read,
			ProxyBindingId:    cmt.Encoder.Encode(model.ProxyBindingId, "proxy_bindings"),
			ProxyBindingBotId: cmt.Encoder.Encode(model.PbBotId, "proxy_binding_bots"),
		})
	}

	response.ProxyBindingBots = make(map[string]ProxyBindingBotResponse, 0)
	for _, pbBot := range pbBots {
		var config interface{}
		json.Unmarshal([]byte(pbBot.Config.String), &config)

		proxyBindingBot := cmt.Encoder.Encode(pbBot.Id, "proxy_binding_bots")
		response.ProxyBindingBots[proxyBindingBot] = ProxyBindingBotResponse{
			Id:            cmt.Encoder.Encode(pbBot.Id, "proxy_binding_bots"),
			Status:        pbBot.Status,
			SessionLength: pbBot.SessionLength,
			Config:        config,
		}
	}

	response.LastId = id

	return response
}
