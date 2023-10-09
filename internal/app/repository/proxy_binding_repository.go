package repository

import (
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type ProxyBindingRepository struct {
	fx.In
	Db *gorm.DB
}

func (pbr ProxyBindingRepository) GetById(id int) models.ProxyBinding {
	model := models.ProxyBinding{}
	pbr.Db.Model(models.ProxyBinding{}).Where(models.ProxyBinding{ID: id}).First(&model)
	return model
}
