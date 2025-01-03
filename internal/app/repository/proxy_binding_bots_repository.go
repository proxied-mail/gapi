package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"time"
)

type ProxyBindingBotsRepositoryInterface interface {
	GetById(id int) (models.ProxyBindingBots, error)
	GetByPbId(pbId int) (models.ProxyBindingBots, error)
	Create(
		botId int,
		pbId int,
		sessionLength int,
		config map[string]interface{},
		demandCc bool,
		allowInterruption bool,
	) models.ProxyBindingBots
	GetByIdIn(ids map[int]int) []models.ProxyBindingBots
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

func (c ProxyBindingBotsRepository) Create(
	botId int,
	pbId int,
	sessionLength int,
	config map[string]interface{},
	demandCc bool,
	allowInterruption bool,
) models.ProxyBindingBots {

	json3, _ := json.Marshal(config)

	model := models.ProxyBindingBots{}
	model.BotId = botId
	model.Status = models.PB_BOT_STATUS_ACTIVE
	model.Config = sql.NullString{String: string(json3), Valid: true}
	model.ProxyBindingId = pbId
	model.SessionLength = sessionLength
	model.DemandCc = demandCc
	model.AllowInterruption = allowInterruption
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()

	c.Db.Save(&model)

	return model
}

func (c ProxyBindingBotsRepository) GetByIdIn(ids map[int]int) []models.ProxyBindingBots {
	var newIds []int
	for _, id := range ids {
		newIds = append(newIds, id)
	}

	var modelsList []models.ProxyBindingBots
	c.Db.Model(models.ProxyBindingBots{}).Where("id IN ?", newIds).Find(&modelsList)

	return modelsList
}
