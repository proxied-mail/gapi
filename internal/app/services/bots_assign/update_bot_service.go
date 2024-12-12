package bots_assign

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/abrouter/gapi/internal/app/http/request/bots_req"
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/access_checker"
	"github.com/abrouter/gapi/pkg/entityId"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type UpdateBotServiceInterface interface {
	UpdateBot(
		user models.UserModel,
		request bots_req.UpdateRequest,
	) (models.ProxyBindingBots, error)
}

type UpdateBotService struct {
	fx.In
	entityId.Encoder
	repository.ProxyBindingRepository
	access_checker.AccessChecker
	repository.BotsRepositoryInterface
	repository.ProxyBindingBotsRepositoryInterface
	*gorm.DB
}

func (abs UpdateBotService) UpdateBot(
	user models.UserModel,
	request bots_req.UpdateRequest,
) (models.ProxyBindingBots, error) {
	model := models.ProxyBindingBots{}

	proxyBindingEncoded, err := abs.Decode(request.ProxyBinding, "proxy_bindings")
	if err != nil {
		return model, errors.New(err.Error() + "cant decode")
	}
	pbModel := abs.ProxyBindingRepository.GetById(int(proxyBindingEncoded))
	if pbModel.Id < 1 {
		return model, errors.New("cant fid pb model")
	}
	if !abs.AccessChecker.CheckProxyBindingAccess(user.Id, pbModel) {
		return model, errors.New("user doesnt have an access to proxy email")
	}

	botId := 0
	if len(request.BotUid) > 0 {
		bot := abs.BotsRepositoryInterface.GetByUid(request.BotUid)
		if bot.Id < 1 {
			return model, errors.New("Cannot find bot by uid")
		}
		botId = bot.Id
	}

	m, e := abs.ProxyBindingBotsRepositoryInterface.GetByPbId(pbModel.Id)
	if e != nil {
		return model, e
	}

	json3, _ := json.Marshal(request.Config)

	m.Status = request.Status
	m.Config = sql.NullString{String: string(json3), Valid: true}
	m.SessionLength = request.SessionLength
	m.DemandCc = request.DemandCc
	m.AllowInterruption = request.AllowInterruption
	m.BotId = botId
	abs.DB.Save(&m)

	return m, nil
}
