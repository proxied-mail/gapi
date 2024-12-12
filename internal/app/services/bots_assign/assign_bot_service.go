package bots_assign

import (
	"errors"
	"github.com/abrouter/gapi/internal/app/http/request/bots_req"
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/access_checker"
	"github.com/abrouter/gapi/pkg/entityId"
	"go.uber.org/fx"
)

type AssignBotServiceInterface interface {
	AssignBot(
		user models.UserModel,
		request bots_req.AssignBotRequest,
	) (models.ProxyBindingBots, error)
}

type AssignBotService struct {
	fx.In
	entityId.Encoder
	repository.ProxyBindingRepository
	access_checker.AccessChecker
	repository.BotsRepositoryInterface
	repository.ProxyBindingBotsRepositoryInterface
}

func (abs AssignBotService) AssignBot(
	user models.UserModel,
	request bots_req.AssignBotRequest,
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

	e, _ := abs.ProxyBindingBotsRepositoryInterface.GetByPbId(pbModel.Id)
	if e.Id > 0 {
		return model, errors.New("bot is already exists")
	}

	newModel := abs.ProxyBindingBotsRepositoryInterface.Create(
		botId,
		pbModel.Id,
		request.SessionLength,
		request.Config,
		request.DemandCc,
		request.AllowInterruption,
	)

	return newModel, nil
}
