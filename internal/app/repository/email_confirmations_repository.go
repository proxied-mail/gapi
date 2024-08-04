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

func (ecr EmailConfirmationsRepository) GetByIdAndUserId(id int, userId int) models.EmailConfirmations {
	model := models.EmailConfirmations{}
	ecr.Db.Model(models.EmailConfirmations{}).Where(models.EmailConfirmations{
		UserId: userId,
		ID:     id,
	}).First(&model)

	return model
}
func (ecr EmailConfirmationsRepository) FirstUnconfirmedNotShown(userId int) models.EmailConfirmations {
	model := models.EmailConfirmations{}
	ecr.Db.Model(models.EmailConfirmations{}).Where(models.EmailConfirmations{
		UserId:                   userId,
		Type:                     2,
		ShownConfirmationRequest: false,
	}).First(&model)

	return model
}
