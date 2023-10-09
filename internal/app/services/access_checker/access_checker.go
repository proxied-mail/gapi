package access_checker

import (
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
)

type AccessChecker struct {
	fx.In
}

func (ac AccessChecker) CheckProxyBindingAccess(userId int, proxyBinding models.ProxyBinding) bool {
	return userId == proxyBinding.UserId
}
