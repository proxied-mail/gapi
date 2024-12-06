package controller

import (
	http2 "github.com/abrouter/gapi/internal/app/http"
	"github.com/abrouter/gapi/internal/app/http/response/conversation_messages_rsp"
	models2 "github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/access_checker"
	"github.com/abrouter/gapi/pkg/entityId"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"net/http"
)

type ConversationsController struct {
	fx.In
	repository.UserRepository
	repository.ProxyBindingRepository
	entityId.Encoder
	repository.ProxyBindingBotMessagesRepositoryInterface
	repository.ProxyBindingBotsRepositoryInterface
	repository.ReceivedEmailsRepositoryInterface
	access_checker.AccessChecker
	conversation_messages_rsp.ConversationMessagesTransformer
	repository.BotsRepositoryInterface
}

func (con ConversationsController) GetMessages(c echo.Context) error {
	lastId := c.QueryParam("lastId")
	lastIdDecoded, err := con.Decode(lastId, "proxy_binding_bot_messages")
	if len(lastId) > 0 && err != nil {
		{
			return c.JSON(http.StatusForbidden, ErrorResponse{Message: "Cant find last value"})
		}
	}

	botUid := c.QueryParam("botUid")

	proxyBindingId := c.QueryParam("proxyBinding")
	var proxyBindingDecoded int64
	if len(proxyBindingId) > 0 {
		var err2 error
		proxyBindingDecoded, err2 = con.Decode(proxyBindingId, "proxy_bindings")
		if err2 != nil {
			return c.JSON(http.StatusNotFound, ErrorResponse{Message: "cant find proxy email"})
		}

		if err2 != nil {
			return c.JSON(http.StatusNotFound, ErrorResponse{Message: "cant find proxy email"})
		}
	}

	if len(proxyBindingId) < 1 && len(botUid) < 1 {
		return c.JSON(http.StatusUnprocessableEntity, ErrorResponse{Message: "Not specified proxy email"})
	}

	currentUser := http2.CurrentUser(c)
	userModel := con.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)

	if userModel.Id < 1 {
		return c.JSON(http.StatusForbidden, ErrorResponse{Message: "Access Denied"})
	}

	var models []models2.ProxyBindingBotMessages
	if proxyBindingDecoded > 0 {
		pb := con.ProxyBindingRepository.GetById(int(proxyBindingDecoded))
		if pb.Id < 1 {
			return c.JSON(http.StatusNotFound, ErrorResponse{Message: "Cant find proxy email"})
		}

		hasPbAccess := con.AccessChecker.CheckProxyBindingAccess(userModel.Id, pb)
		if !hasPbAccess {
			return c.JSON(http.StatusForbidden, ErrorResponse{Message: "Access denied to the proxy email"})
		}

		pbBot, err3 := con.ProxyBindingBotsRepositoryInterface.GetByPbId(pb.Id)
		if err3 != nil {
			return c.JSON(http.StatusNotFound, ErrorResponse{Message: "Bot is not found"})
		}

		models = con.ProxyBindingBotMessagesRepositoryInterface.Query(pbBot.Id, int(lastIdDecoded))
	} else {

		bot := con.BotsRepositoryInterface.GetByUid(botUid)
		if bot.UserId != userModel.Id {
			return c.JSON(http.StatusForbidden, ErrorResponse{Message: "Dont have an access to the bot"})
		}

		models = con.ProxyBindingBotMessagesRepositoryInterface.QueryByBotUid(botUid, int(lastIdDecoded))
	}

	var receivedEmailsIds []int
	for _, model := range models {
		receivedEmailsIds = append(receivedEmailsIds, model.ReceivedEmailId)
	}

	receivedEmails, _ := con.ReceivedEmailsRepositoryInterface.GetIn(receivedEmailsIds)
	rsp := con.ConversationMessagesTransformer.Transform(models, receivedEmails)

	return c.JSON(http.StatusOK, rsp)
}
