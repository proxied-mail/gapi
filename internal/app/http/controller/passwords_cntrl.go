package controller

import (
	json2 "encoding/json"
	http2 "github.com/abrouter/gapi/internal/app/http"
	real_emails "github.com/abrouter/gapi/internal/app/http/request/passwords"
	"github.com/abrouter/gapi/internal/app/http/response/common"
	"github.com/abrouter/gapi/internal/app/http/response/passwords_rsp"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/entity_fetcher"
	"github.com/abrouter/gapi/internal/app/services/password_srv"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	validator2 "gopkg.in/go-playground/validator.v9"
	"io"
	"net/http"
)

type PasswordsCntrl struct {
	fx.In
	repository.UserRepository
	ProxyBindingFetcher entity_fetcher.ProxyBindingFetcher
	password_srv.PasswordUpdater
	passwords_rsp.PasswordListResponseMapper
}

func (cntrl PasswordsCntrl) Update(c echo.Context) error {
	currentUser := http2.CurrentUser(c)
	userModel := cntrl.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)

	request := real_emails.ProxyBindingPasswordUpdate{}
	reqBody, err := io.ReadAll(c.Request().Body)

	if err != nil {
		resp, _ := json2.Marshal(ErrorResponse{
			Message: "Invalid json",
			Status:  false,
		})
		return c.String(http.StatusUnprocessableEntity, string(resp))
	}
	json2.Unmarshal(reqBody, &request)
	valid := validator2.New().Struct(request)
	if valid != nil {
		resp, _ := json2.Marshal(ErrorResponse{
			Message: valid.Error(),
			Status:  false,
		})
		return c.String(http.StatusUnprocessableEntity, string(resp))
	}
	proxyBindingModel, err2 := cntrl.ProxyBindingFetcher.GetModel(request.ProxyBindingId)
	if err2 != nil || proxyBindingModel.Id == 0 {
		resp, _ := json2.Marshal(ErrorResponse{
			Message: "Cannot find proxy binding entity",
			Status:  false,
		})
		return c.String(http.StatusUnprocessableEntity, string(resp))
	}

	_, err3 := cntrl.PasswordUpdater.UpdatePasswordByProxyBinding(userModel, proxyBindingModel, request.Password)
	if err3 != nil {
		resp, _ := json2.Marshal(ErrorResponse{
			Message: "Cannot update model",
			Status:  false,
		})
		return c.String(http.StatusUnprocessableEntity, string(resp))
	}

	resp := common.GetSuccess()
	return c.String(http.StatusOK, resp)
}

func (cntrl PasswordsCntrl) List(c echo.Context) error {
	currentUser := http2.CurrentUser(c)
	userModel := cntrl.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)
	list := cntrl.PasswordsRepository.AllByUser(userModel.Id)
	rsp := cntrl.PasswordListResponseMapper.MapResponse(list)
	rspJson, _ := json2.Marshal(rsp)

	return c.String(http.StatusOK, string(rspJson))
}
