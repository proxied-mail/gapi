package passwords_rsp

import (
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/pkg/entityId"
	"go.uber.org/fx"
	"time"
)

type PasswordRsp struct {
	Id            string    `json:"id"`
	UserId        string    `json:"user_id"`
	RelatedToType int       `json:"related_to_type"`
	RelatedToId   string    `json:"related_to_id"`
	Login         string    `json:"login"`
	Website       string    `json:"website"`
	Password      string    `json:"password"`
	Title         string    `json:"title"`
	Note          string    `json:"note"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type PasswordListResponseMapper struct {
	fx.In
	entityId.Encoder
}

func (plrm PasswordListResponseMapper) MapResponse(passwords []models.Passwords) []PasswordRsp {

	var newModels = make([]PasswordRsp, 0)
	for _, model := range passwords {
		newModels = append(newModels, PasswordRsp{
			Id:            plrm.Encode(model.ID, "passwords"),
			UserId:        plrm.Encode(model.UserId, "users"),
			RelatedToType: model.RelatedToType,
			RelatedToId:   plrm.Encode(model.RelatedToId, "proxy_binding"), //todo more logic here
			Password:      model.Password,
			Login:         model.Login.String,
			Title:         model.Title.String,
			Website:       model.Website.String,
			Note:          model.Note.String,
			CreatedAt:     model.CreatedAt,
			UpdatedAt:     model.UpdatedAt,
		})
	}

	return newModels
}
