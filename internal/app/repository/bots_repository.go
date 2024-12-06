package repository

import (
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type BotsRepositoryInterface interface {
	GetByUid(uid string) models.Bots
}

type BotsRepository struct {
	fx.In
	*gorm.DB
}

func (br BotsRepository) GetByUid(uid string) models.Bots {
	r := models.Bots{}
	br.DB.Model(models.Bots{}).Where("uid = ?", uid).First(&r)

	return r
}
