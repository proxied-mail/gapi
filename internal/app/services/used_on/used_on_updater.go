package used_on

import (
	"errors"
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/internal/app/services/access_checker"
	"go.uber.org/fx"
)

type UsedOnUpdaterStruct struct {
	fx.In
	access_checker.AccessChecker
}

func (uo UsedOnUpdaterStruct) Update(
	user models.UserModel,
	proxyBinding models.ProxyBinding,
	list []string) error {
	if uo.AccessChecker.CheckProxyBindingAccess(user.Id, proxyBinding) {
		return errors.New("user id and proxy binding user id not equal")
	}

}
