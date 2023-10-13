package used_on

type UsedOnRequest struct {
	ProxyBindingId string   `json:"proxy_binding_id" validate:"required"`
	List           []string `json:"list" validate:"required"`
}
