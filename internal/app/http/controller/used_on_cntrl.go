package controller

import (
	json2 "encoding/json"
	http2 "github.com/abrouter/gapi/internal/app/http"
	"github.com/abrouter/gapi/internal/app/http/request/used_on"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/entity_fetcher"
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

}
