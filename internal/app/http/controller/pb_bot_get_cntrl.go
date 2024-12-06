package controller

import (
	http2 "github.com/abrouter/gapi/internal/app/http"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/access_checker"
	"github.com/abrouter/gapi/internal/app/services/bots_assign"
	"github.com/abrouter/gapi/pkg/entityId"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"net/http"
)

type PbBotGetController struct {
	fx.In
	bots_assign.AssignBotServiceInterface
	repository.UserRepository
	access_checker.AccessChecker
	entityId.Encoder
	repository.ProxyBindingRepository
	repository.ProxyBindingBotsRepositoryInterface
}

type PbBotResponse struct {
	Items []PbBotResponseItem `json:"items"`
}

type PbBotResponseItem struct {
	Status           int `json:"status"`
	SessionLength    int `json:"session_length"`
	MessagesReceived int `json:"messages_received"`
	MessagesSent     int `json:"messages_sent"`
}

func (con PbBotGetController) Get(c echo.Context) error {
	proxyBindingId := c.QueryParam("proxyBinding")
	proxyBindingDecoded, err2 := con.Decode(proxyBindingId, "proxy_bindings")
	if err2 != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: "cant find proxy email"})
	}
	proxyBindingDecodedInt := int(proxyBindingDecoded)

	if err2 != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: "cant find proxy email"})
	}

	currentUser := http2.CurrentUser(c)
	userModel := con.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)

	if userModel.Id < 1 {
		return c.JSON(http.StatusForbidden, ErrorResponse{Message: "Access Denied"})
	}

	pb := con.ProxyBindingRepository.GetById(proxyBindingDecodedInt)
	if pb.Id < 1 {
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: "Pb not found"})
	}
	if !con.AccessChecker.CheckProxyBindingAccess(userModel.Id, pb) {
		return c.JSON(http.StatusForbidden, ErrorResponse{Message: "Access Denied"})
	}

	m, _ := con.GetByPbId(pb.Id)
	rsp := PbBotResponse{}
	if m.Id > 0 {
		rsp.Items = append(rsp.Items, PbBotResponseItem{
			Status:           m.Status,
			SessionLength:    m.SessionLength,
			MessagesReceived: 0,
			MessagesSent:     0,
		})
	}

	return c.JSON(http.StatusOK, rsp)
}
