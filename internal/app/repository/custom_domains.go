package repository

import (
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type CustomDomainsRepository struct {
	fx.In
	Db *gorm.DB
}

func (cdr CustomDomainsRepository) GetAllByUser(UserId int) []models.CustomDomain {
	var list []models.CustomDomain
	cdr.Db.Model(models.CustomDomain{}).Where(models.CustomDomain{UserId: UserId}).Find(&list)
	return list
}
