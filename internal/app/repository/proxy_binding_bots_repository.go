package repository

import (
	"errors"
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type ProxyBindingBotsRepositoryInterface interface {
	GetById(id int) (models.ProxyBindingBots, error)
}

type ProxyBindingBotsRepository struct {
	fx.In
	Db *gorm.DB
}

func (c ProxyBindingBotsRepository) GetById(id int) (models.ProxyBindingBots, error) {
	var model models.ProxyBindingBots
	c.Db.Model(models.ProxyBindingBots{}).Where("id", id).First(&model)

	if model.Id < 1 {
		return model, errors.New("Failed to find pb bot")
	}

	return model, nil
}
