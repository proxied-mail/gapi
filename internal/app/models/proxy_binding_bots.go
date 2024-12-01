package models

type ProxyBindingBots struct {
	Id               int `json:"id"`
	ProxyBindingId   int `json:"proxy_binding_id"`
	Status           int `json:"status"`
	SessionLength    int `json:"session_length"`
	MessagesReceived int `json:"messages_received"`
	MessagesSent     int `json:"messages_sent"`
}

func (pbb ProxyBindingBots) GetTableName() string {
	return "proxy_binding_bots"
}
