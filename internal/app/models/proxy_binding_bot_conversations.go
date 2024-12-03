package models

import "time"

type ProxyBindingBotConversations struct {
	Id                    int       `json:"id"`
	PbBotId               int       `json:"pb_bot_id"`
	ProxyBindingId        int       `json:"proxy_binding_id"`
	SenderEmail           string    `json:"sender_email"`
	LastMessageAt         time.Time `json:"last_message_at"`
	ReceivedMessagesCount int       `json:"received_messages_count"`
	SentMessagesCount     int       `json:"sent_messages_count"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
