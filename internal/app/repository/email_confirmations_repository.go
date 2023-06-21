package repository

import (
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type EmailConfirmationsRepository struct {
	fx.In
	Db *gorm.DB
}

func (ecr EmailConfirmationsRepository) GetAllConfirmedEmails(userId int) []string {
	var list []models.EmailConfirmations
	ecr.Db.Model(models.EmailConfirmations{}).Where(models.EmailConfirmations{
		UserId:    userId,
		Confirmed: 1,
	}).Find(&list)
	var emailsList []string
	var model models.EmailConfirmations

	for _, model = range list {
		emailsList = append(emailsList, model.Email)
	}
	return emailsList
}
