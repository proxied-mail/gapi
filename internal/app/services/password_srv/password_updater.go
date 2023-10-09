package password_srv

import (
	"errors"
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/access_checker"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type PasswordUpdater struct {
	fx.In
	access_checker.AccessChecker
	repository.PasswordsRepository
	*gorm.DB
}

func (upd PasswordUpdater) UpdatePasswordByProxyBinding(
	user models.UserModel,
	proxyBinding models.ProxyBinding,
	Password string,
) (models.Passwords, error) {
	hasAccess := upd.AccessChecker.CheckProxyBindingAccess(user.Id, proxyBinding)

	if !hasAccess {
		return models.Passwords{}, errors.New("access denied")
	}

	current := upd.PasswordsRepository.GetPasswordByProxyBinding(proxyBinding.Id, user.Id)
	if current.ID > 0 {
		current.Password = Password
		upd.Db.Save(&current)
		return current, nil
	}

	newModel := models.Passwords{}
	newModel.UserId = user.Id
	newModel.Password = Password
	newModel.RelatedToId = proxyBinding.Id
	newModel.RelatedToType = models.RELATED_TO_TYPE_PROXY_BINDING
	upd.Db.Create(&newModel)
	return newModel, nil
}
