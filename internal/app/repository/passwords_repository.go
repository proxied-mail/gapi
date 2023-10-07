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
	pr.Db.Model(models.Passwords{}).Where(models.Passwords{})
}
