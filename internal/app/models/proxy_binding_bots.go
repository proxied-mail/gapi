package models

import "time"

type ProxyBindingBots struct {
	Id               int       `json:"id"`
	BotId            int       `json:"bot_id"`
	ProxyBindingId   int       `json:"proxy_binding_id"`
	Status           int       `json:"status"`
	SessionLength    int       `json:"session_length"`
	MessagesReceived int       `json:"messages_received"`
	MessagesSent     int       `json:"messages_sent"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (pbb ProxyBindingBots) GetTableName() string {
	return "proxy_binding_bots"
}
