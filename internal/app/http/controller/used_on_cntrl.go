package controller

import (
	json2 "encoding/json"
	http2 "github.com/abrouter/gapi/internal/app/http"
	"github.com/abrouter/gapi/internal/app/http/request/used_on"
	"github.com/abrouter/gapi/internal/app/http/response/common"
	"github.com/abrouter/gapi/internal/app/http/response/used_on_rsp"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/entity_fetcher"
	used_on2 "github.com/abrouter/gapi/internal/app/services/used_on"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	validator2 "gopkg.in/go-playground/validator.v9"
	"io"
	"net/http"
)

type UsedOnCntrl struct {
	fx.In
	repository.UserRepository
	entity_fetcher.ProxyBindingFetcher
	used_on2.UsedOnUpdater
	used_on_rsp.UsedOnResponse
}

func (cntrl UsedOnCntrl) Change(c echo.Context) error {
	currentUser := http2.CurrentUser(c)
	userModel := cntrl.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)

	req := used_on.UsedOnRequest{}
	reqBody, err := io.ReadAll(c.Request().Body)
	if err != nil {
		resp, _ := json2.Marshal(ErrorResponse{
			Message: "Cannot parse request",
			Status:  false,
		})
		return c.String(http.StatusUnprocessableEntity, string(resp))
	}

	json2.Unmarshal(reqBody, &req)

	valid := validator2.New().Struct(req)
	if valid != nil {
		resp, _ := json2.Marshal(ErrorResponse{
			Message: valid.Error(),
			Status:  false,
		})
		return c.String(http.StatusUnprocessableEntity, string(resp))
	}
	proxyBindingModel, err2 := cntrl.ProxyBindingFetcher.GetModel(req.ProxyBindingId)
	if err2 != nil || proxyBindingModel.Id == 0 {
		resp, _ := json2.Marshal(ErrorResponse{
			Message: "Cannot find proxy binding entity",
			Status:  false,
		})
		return c.String(http.StatusUnprocessableEntity, string(resp))
	}

	err3 := cntrl.UsedOnUpdater.Update(userModel, proxyBindingModel, req.List)
	if err3 != nil {
		resp, _ := json2.Marshal(ErrorResponse{
			Message: "Cannot update used on entity:" + err3.Error(),
			Status:  false,
		})
		return c.String(http.StatusUnprocessableEntity, string(resp))
	}

	resp := common.GetSuccess()
	return c.String(http.StatusOK, resp)
}

func (cntrl UsedOnCntrl) List(c echo.Context) error {
	currentUser := http2.CurrentUser(c)
	userModel := cntrl.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)

	all := cntrl.UsedOnRepository.GetUsedOnByUserId(userModel.Id)
	rspModels := cntrl.UsedOnResponse.MapResponse(all)

	rsp, _ := json2.Marshal(rspModels)
	return c.String(http.StatusOK, string(rsp))
}
