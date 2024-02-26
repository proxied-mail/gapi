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

func (cdr CustomDomainsRepository) UserHasDomain(userId int, domain string) bool {
	var count int64
	cdr.Db.Model(models.CustomDomain{}).Where(models.CustomDomain{
		UserId: userId,
		Domain: domain,
	}).Count(&count)

	return count > 0
}

func (cdr CustomDomainsRepository) GetAllAvailable(userId int) []models.CustomDomain {
	var list []models.CustomDomain
	cdr.Db.Model(models.CustomDomain{}).Where(models.CustomDomain{
		UserId: userId,
	}).Where("(status) IN (4,5)").Or(models.CustomDomain{
		IsShared: true,
	}).Find(&list)
	return list
}
