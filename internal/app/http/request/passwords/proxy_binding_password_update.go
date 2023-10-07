package real_emails

type ProxyBindingPasswordUpdate struct {
	ProxyBindingId string `json:"proxy_binding_id" validate:"required"`
	Password       string `json:"password" validate:"required"`
}
