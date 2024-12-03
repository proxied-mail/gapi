package repository

import (
	"errors"
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type ReceivedEmailsRepositoryInterface interface {
	GetOneById(id int) (models.ReceivedEmails, error)
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
