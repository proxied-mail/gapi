package repository

import (
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type UsedOnRepository struct {
	fx.In
	*gorm.DB
}

func (r UsedOnRepository) GetUsedOnByProxyBindingId(proxyBindingId int) models.ProxyBindingUsedOn {
	var model models.ProxyBindingUsedOn
	r.DB.Model(models.ProxyBindingUsedOn{}).Where(models.ProxyBindingUsedOn{ProxyBindingId: proxyBindingId}).First(
		&model,
	)

	return model
}

func (r UsedOnRepository) GetUsedOnByUserId(userId int) []models.ProxyBindingUsedOn {
	var modelsList []models.ProxyBindingUsedOn
	r.DB.Model(models.ProxyBindingUsedOn{}).Where(models.ProxyBindingUsedOn{UserId: userId}).Find(
		&modelsList,
	)

	return modelsList
}
