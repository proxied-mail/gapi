package repository

import (
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type SettingsRepository struct {
	fx.In
	Db *gorm.DB
}

func (sr SettingsRepository) AllUserSettings(UserId int) []models.Settings {
	var modelsList []models.Settings
	sr.Db.Model(models.Settings{}).Where(models.Settings{UserId: UserId}).Find(
		&modelsList,
	)

	return modelsList
}
