package settings

import (
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/internal/app/repository"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type UpdateSettingsRequest struct {
	Settings []SettingRequest `json:"settings"`
}

type SettingRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type UpdateSettingsService struct {
	fx.In
	SettingsRepository repository.SettingsRepository
	Db                 *gorm.DB
}

func (uss UpdateSettingsService) UpdateSettings(
	user models.UserModel,
	request UpdateSettingsRequest,
) ([]models.Settings, error) {
	userSettings := uss.SettingsRepository.AllUserSettings(user.Id)

	var newModels []models.Settings
	for _, setting := range request.Settings {
		isSet := false
		for _, userSetting := range userSettings {
			if setting.Key == userSetting.Name {
				userSetting.Value = setting.Value
				uss.Db.Save(&userSetting)
				newModels = append(newModels, userSetting)
				isSet = true
			}
		}
		if isSet {
			continue
		}

		newSetting := models.Settings{
			UserId: user.Id,
			Name:   setting.Key,
			Value:  setting.Value,
		}
		uss.Db.Create(&newSetting)
		newModels = append(newModels, newSetting)
	}

	return newModels, nil
}
