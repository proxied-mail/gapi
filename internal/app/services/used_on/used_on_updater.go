package used_on

import (
	"encoding/json"
	"errors"
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/access_checker"
	"go.uber.org/fx"
)

type UsedOnUpdater struct {
	fx.In
	access_checker.AccessChecker
	repository.UsedOnRepository
}

func (uo UsedOnUpdater) Update(
	user models.UserModel,
	proxyBinding models.ProxyBinding,
	list []string) error {
	if !uo.AccessChecker.CheckProxyBindingAccess(user.Id, proxyBinding) {
		return errors.New("user id and proxy binding user id not equal")
	}

	listJson, _ := json.Marshal(list)
	usedOn := uo.UsedOnRepository.GetUsedOnByProxyBindingId(proxyBinding.Id)
	if usedOn.Id > 0 {
		usedOn.JsonList = string(listJson)
		uo.DB.Save(&usedOn)
		return nil
	}

	usedOnNew := models.ProxyBindingUsedOn{}
	usedOnNew.ProxyBindingId = proxyBinding.Id
	usedOnNew.UserId = user.Id
	usedOnNew.JsonList = string(listJson)
	uo.DB.Create(&usedOnNew)

	return nil
}
