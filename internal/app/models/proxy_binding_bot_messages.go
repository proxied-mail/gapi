package models

import (
	"time"
)

type ProxyBindingBotMessages struct {
	Id              int       `json:"id"`
	PbBotId         int       `json:"pb_bot_id"`
	ProxyBindingId  int       `json:"proxy_binding_id"`
	SenderEmail     string    `json:"sender_email"`
	ReceivedEmailId int       `json:"received_email_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
