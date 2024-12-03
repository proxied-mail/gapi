package controller

import (
	"encoding/json"
	"github.com/abrouter/gapi/internal/app/http/request/bots_req"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/bot_messages"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"io"
	"net/http"
)

type BotController struct {
	fx.In
	repository.ProxyBindingBotMessagesRepositoryInterface
	bot_messages.MessageSaverServiceInterface
}

type ReceivedEmailNotifyResponse struct {
	Status bool `json:"status"`
}

func (bc BotController) ReceivedEmailNotify(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		resp := ErrorResponse{
			Message: "Cant parse request",
			Status:  false,
		}
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	req := bots_req.BotsRequestPbNotifyReceivedEmail{}
	validationErr := json.Unmarshal(body, &req)
	if validationErr != nil {
		resp := ErrorResponse{
			Message: "validation error:" + validationErr.Error(),
			Status:  false,
		}
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	dto := bot_messages.MessageSaverServiceDto{
		ReceivedEmailId:   req.ReceivedEmailId,
		ProxyBindingBotId: req.ProxyBindingBotId,
	}
	creatingError := bc.MessageSaverServiceInterface.Save(dto)

	if creatingError != nil {
		resp := ErrorResponse{
			Message: "unable to create entity in db:" + creatingError.Error(),
			Status:  false,
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	rsp2 := ReceivedEmailNotifyResponse{
		Status: true,
	}

	return c.JSON(http.StatusCreated, rsp2)
}
