package controller

import (
	"encoding/json"
	http2 "github.com/abrouter/gapi/internal/app/http"
	domains2 "github.com/abrouter/gapi/internal/app/http/response/domains"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/domains"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type DomainsController struct {
	fx.In
	UserRepository      repository.UserRepository
	CreateDomainService domains.CreateDomainService
	repository.CustomDomainsRepository
	StatusProcessorService domains.StatusProcessorService
}

func (dc DomainsController) Create(
	c echo.Context,
) error {
	currentUser := http2.CurrentUser(c)
	userModel := dc.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)
	createDomainRequest := domains.CreateDomainRequest{}
	json.NewDecoder(c.Request().Body).Decode(&createDomainRequest)

	model, err := dc.CreateDomainService.CreateDomain(
		userModel.Id,
		createDomainRequest,
	)
	if err != nil {
		resp, _ := json.Marshal(ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
		return c.String(http.StatusUnprocessableEntity, string(resp))
	}
	model2 := domains2.MapResponse(model)

	resp, _ := json.Marshal(model2)
	return c.String(http.StatusOK, string(resp))
}

func (dc DomainsController) List(
	c echo.Context,
) error {
	currentUser := http2.CurrentUser(c)
	userModel := dc.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)
	list := dc.CustomDomainsRepository.GetAllAvailable(userModel.Id)
	mappedList := domains2.MapResponseList(list)
	resp, _ := json.Marshal(mappedList)
	return c.String(http.StatusOK, string(resp))
}

func (dc DomainsController) ListCustom(
	c echo.Context,
) error {
	currentUser := http2.CurrentUser(c)
	userModel := dc.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)

	list := dc.CustomDomainsRepository.GetAllByUser(userModel.Id)
	mappedList := domains2.MapResponseList(list)
	if c.FormValue("ignoreProcessing") != "1" {
		mappedList = dc.StatusProcessorService.ProcessStatus(mappedList)
	}

	resp, _ := json.Marshal(mappedList)
	return c.String(http.StatusOK, string(resp))
}
