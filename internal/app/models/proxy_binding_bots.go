package models

import (
	"database/sql"
	"time"
)

const PB_BOT_STATUS_ACTIVE = 3
const PB_BOT_STATUS_UNACTIVE = 2

type ProxyBindingBots struct {
	Id                int            `json:"id"`
	BotId             int            `json:"bot_id"`
	ProxyBindingId    int            `json:"proxy_binding_id"`
	Status            int            `json:"status"`
	Config            sql.NullString `json:"config"`
	SessionLength     int            `json:"session_length"`
	AllowInterruption bool           `json:"allow_interruption"`
	DemandCc          bool           `json:"demand_cc"`
	MessagesReceived  int            `json:"messages_received"`
	MessagesSent      int            `json:"messages_sent"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
}

func (pbb ProxyBindingBots) GetTableName() string {
	return "proxy_binding_bots"
}
