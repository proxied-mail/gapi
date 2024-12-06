package controller

import (
	"encoding/json"
	http2 "github.com/abrouter/gapi/internal/app/http"
	"github.com/abrouter/gapi/internal/app/http/request/bots_req"
	"github.com/abrouter/gapi/internal/app/http/response/common"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/bots_assign"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"io"
	"net/http"
)

type UpdateBotController struct {
	fx.In
	bots_assign.UpdateBotServiceInterface
	repository.UserRepository
}

func (abc UpdateBotController) UpdateBot(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ErrorResponse{Message: "Cant process payload"})
	}
	r := bots_req.UpdateRequest{}
	err2 := json.Unmarshal(b, &r)
	if err2 != nil {
		return c.JSON(http.StatusUnprocessableEntity, ErrorResponse{Message: "Cant process payload 2"})
	}

	currentUser := http2.CurrentUser(c)
	userModel := abc.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)

	_, err3 := abc.UpdateBotServiceInterface.UpdateBot(userModel, r)
	if err3 != nil {
		return c.JSON(http.StatusUnprocessableEntity, ErrorResponse{
			Message: "Cant update bot, error" + err3.Error(),
		})
	}

	return c.JSON(http.StatusCreated, common.SuccessWithMsg{Message: "Successfully updated", Status: true})
}
