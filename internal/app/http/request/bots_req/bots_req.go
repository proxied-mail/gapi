package bots_req

type BotsRequestPbNotifyReceivedEmail struct {
	ProxyBindingBotId int `json:"proxyBindingBotId" validate:"required"`
	ReceivedEmailId   int `json:"receivedEmailId" validate:"required"`
}

type AssignBotRequest struct {
	BotUid        string `json:"bot_uid"`
	ProxyBinding  string `json:"proxy_binding_id" validate:"required"`
	SessionLength int    `json:"session_length" validate:"required"`
}
