package repository

import (
	"errors"
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"time"
)

type ProxyBindingBotsRepositoryInterface interface {
	GetById(id int) (models.ProxyBindingBots, error)
	GetByPbId(pbId int) (models.ProxyBindingBots, error)
	Create(botId int, pbId int, sessionLength int) models.ProxyBindingBots
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

func (c ProxyBindingBotsRepository) GetByPbId(pbId int) (models.ProxyBindingBots, error) {
	var model models.ProxyBindingBots
	c.Db.Model(models.ProxyBindingBots{}).Where("proxy_binding_id", pbId).First(&model)

	if model.Id < 1 {
		return model, errors.New("Failed to find pb bot")
	}

	return model, nil
}

func (c ProxyBindingBotsRepository) Create(botId int, pbId int, sessionLength int) models.ProxyBindingBots {
	model := models.ProxyBindingBots{}
	model.BotId = botId
	model.Status = 3
	model.ProxyBindingId = pbId
	model.SessionLength = sessionLength
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()

	c.Db.Save(&model)

	return model
}
