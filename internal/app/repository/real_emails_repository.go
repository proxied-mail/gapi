package repository

import (
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type RealEmailsRepository struct {
	fx.In
	Db *gorm.DB
}

func (rer RealEmailsRepository) GetAllUniqueByUser(userId int) []models.RealAddress {
	var list []models.RealAddress

	rer.Db.Model(models.RealAddress{}).Where(models.RealAddress{
		UserId: userId,
	}).Where("deleted_at is null").Group("real_address").Find(&list)
	return list
}
