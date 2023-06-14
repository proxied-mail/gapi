package repository

import (
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func UseUserRepository(db *gorm.DB) *UserRepository {
	ur := UserRepository{
		Gorm: db,
	}
	return &ur
}

type UserRepository struct {
	fx.In
	Gorm *gorm.DB
}

func (r UserRepository) GetUserByEmail(username string) models.UserModel {
	model := models.UserModel{}

	r.Gorm.Where("username", username).First(&model)
	r.Gorm.Model(model).First(&model)
	return model
}
