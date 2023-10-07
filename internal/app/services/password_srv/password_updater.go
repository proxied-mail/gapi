package password_srv

import (
	"errors"
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/internal/app/services/access_checker"
	"go.uber.org/fx"
)

type PasswordUpdater struct {
	fx.In
	AccessChecker access_checker.AccessChecker
}

func (upd PasswordUpdater) updateByProxyBinding(
	user models.UserModel,
	proxyBinding models.ProxyBinding,
	Password string,
) (models.Passwords, error) {
	hasAccess := upd.AccessChecker.CheckProxyBindingAccess(user.Id, proxyBinding)

	if !hasAccess {
		return models.Passwords{}, errors.New("access denied")
	}

}
