package bots_req

type BotsRequestPbNotifyReceivedEmail struct {
	ProxyBindingBotId int `json:"proxyBindingBotId" validate:"required"`
	ReceivedEmailId   int `json:"receivedEmailId" validate:"required"`
}
