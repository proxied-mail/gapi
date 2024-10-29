package controller

import (
	"encoding/json"
	http2 "github.com/abrouter/gapi/internal/app/http"
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/settings"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"net/http"
)

type SettingsController struct {
	fx.In
	UserRepository        repository.UserRepository
	UpdateSettingsService settings.UpdateSettingsService
}

type SettingsResponse struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt string `json:"created_at"`
}

func (sc SettingsController) Update(
	c echo.Context,
) error {
	currentUser := http2.CurrentUser(c)
	userModel := sc.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)
	updateSettingsRequest := settings.UpdateSettingsRequest{}
	json.NewDecoder(c.Request().Body).Decode(&updateSettingsRequest)

	model, err := sc.UpdateSettingsService.UpdateSettings(
		userModel,
		updateSettingsRequest,
	)

	if err != nil {
		resp, _ := json.Marshal(ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
		return c.String(http.StatusUnprocessableEntity, string(resp))
	}
	rsp1 := MapResponse(model)

	resp, _ := json.Marshal(rsp1)
	return c.String(http.StatusOK, string(resp))
}

func MapResponse(model []models.Settings) []SettingsResponse {
	var response []SettingsResponse
	for _, setting := range model {
		response = append(response, SettingsResponse{
			Key:   setting.Name,
			Value: setting.Value,
		})
	}
	return response
}
