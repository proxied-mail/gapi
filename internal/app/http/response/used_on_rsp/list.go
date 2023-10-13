package used_on_rsp

import (
	"encoding/json"
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/pkg/entityId"
	"go.uber.org/fx"
)

type UsedOnList struct {
	ProxyBindingId string   `json:"proxy_binding_id"`
	List           []string `json:"list"`
}

type UsedOnResponse struct {
	fx.In
	entityId.Encoder
}

func (onr UsedOnResponse) MapResponse(models []models.ProxyBindingUsedOn) []UsedOnList {
	list := make([]UsedOnList, 0)
	for _, model := range models {

		var unpacked []string
		_ = json.Unmarshal([]byte(model.JsonList), &unpacked)

		list = append(list, UsedOnList{
			ProxyBindingId: onr.Encode(model.ProxyBindingId, "proxy_bindings"),
			List:           unpacked,
		})
	}

	return list
}
