package controller

import (
	"encoding/json"
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
	Status           int         `json:"status"`
	SessionLength    int         `json:"session_length"`
	Config           interface{} `json:"config"`
	MessagesReceived int         `json:"messages_received"`
	MessagesSent     int         `json:"messages_sent"`
	ExtendsId        string      `json:"extends_id"`
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

		var jsonConfig interface{}
		json.Unmarshal([]byte(m.Config.String), &jsonConfig)

		rsp.Items = append(rsp.Items, PbBotResponseItem{
			Status:           m.Status,
			Config:           jsonConfig,
			SessionLength:    m.SessionLength,
			MessagesReceived: m.MessagesReceived,
			MessagesSent:     m.MessagesSent,
			ExtendsId:        con.Encoder.Encode(m.BotId, "bots"),
		})
	}

	return c.JSON(http.StatusOK, rsp)
}
