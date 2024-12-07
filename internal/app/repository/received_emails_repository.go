package repository

import (
	"errors"
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type ReceivedEmailsRepositoryInterface interface {
	GetOneById(id int) (models.ReceivedEmails, error)
	GetIn(ids []int) (map[int]models.ReceivedEmails, error)
}

type ReceivedEmailsRepository struct {
	fx.In
	Db *gorm.DB
}

func (r ReceivedEmailsRepository) GetOneById(id int) (models.ReceivedEmails, error) {
	model := models.ReceivedEmails{}
	r.Db.Model(model).Where("id", id).First(&model)
	if model.Id < 1 {
		return model, errors.New("cant find model")
	}

	return model, nil
}

func (r ReceivedEmailsRepository) GetIn(ids []int) (map[int]models.ReceivedEmails, error) {
	var result []models.ReceivedEmails
	//ids to []int
	newIds := make([]int, 0)
	for _, id := range ids {
		newIds = append(newIds, id)
	}

	r.Db.Model(models.ReceivedEmails{}).Where("id IN (?)", newIds).Find(&result)
	newModels := make(map[int]models.ReceivedEmails, 0)

	for _, model := range result {
		newModels[model.Id] = model
	}

	return newModels, nil
}
