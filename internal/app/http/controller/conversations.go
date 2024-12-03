package controller

import (
	http2 "github.com/abrouter/gapi/internal/app/http"
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
	access_checker.AccessChecker
}

func (con ConversationsController) GetMessages(c echo.Context) error {
	lastId := c.QueryParam("lastId")
	lastIdDecoded, err := con.Decode(lastId, "proxy_binding_bot_messages")
	if len(lastId) > 0 && err != nil {
		{
			return c.JSON(http.StatusForbidden, ErrorResponse{Message: "Cant find last value"})
		}
	}

	proxyBindingId := c.QueryParam("proxyBinding")
	proxyBindingDecoded, err2 := con.Decode(proxyBindingId, "proxy_bindings")
	if err2 != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: "cant find proxy email"})
	}

	if len(proxyBindingId) < 1 {
		return c.JSON(http.StatusUnprocessableEntity, ErrorResponse{Message: "Not specified proxy email"})
	}

	currentUser := http2.CurrentUser(c)
	userModel := con.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)

	if userModel.Id < 1 {
		return c.JSON(http.StatusForbidden, ErrorResponse{Message: "Access Denied"})
	}

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

	con.ProxyBindingBotMessagesRepositoryInterface.Query(pbBot.Id, int(lastIdDecoded))

	return nil
}
