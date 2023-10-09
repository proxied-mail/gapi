package repository

import (
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type PasswordsRepository struct {
	fx.In
	Db *gorm.DB
}

func (pr PasswordsRepository) GetPasswordByProxyBinding(proxyBindingId int, userId int) models.Passwords {
	model := models.Passwords{}
	pr.Db.Model(models.Passwords{}).Where(models.Passwords{
		RelatedToId:   proxyBindingId,
		RelatedToType: models.RELATED_TO_TYPE_PROXY_BINDING,
		UserId:        userId,
	}).First(&model)
	return model
}

func (pr PasswordsRepository) AllByUser(userId int) []models.Passwords {
	var modelsList []models.Passwords
	pr.Db.Model(models.Passwords{}).Where(models.Passwords{UserId: userId}).Find(&modelsList)
	return modelsList
}
