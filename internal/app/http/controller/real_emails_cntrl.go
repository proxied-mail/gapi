package controller

import (
	json2 "encoding/json"
	http2 "github.com/abrouter/gapi/internal/app/http"
	"github.com/abrouter/gapi/internal/app/http/request/real_emails"
	"github.com/abrouter/gapi/internal/app/http/response/common"
	"github.com/abrouter/gapi/internal/app/http/response/email_confirmations"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/real_emails_srv"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	validator2 "gopkg.in/go-playground/validator.v9"
	"io"
	"net/http"
)

type RealEmailsCntrl struct {
	fx.In
	UserRepository               repository.UserRepository
	EmailConfirmationsRepository repository.EmailConfirmationsRepository
	ReplaceRealEmail             real_emails_srv.ReplaceRealEmail
}

func (rec RealEmailsCntrl) Get(c echo.Context) error {
	currentUser := http2.CurrentUser(c)
	userModel := rec.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)
	emailsList := rec.EmailConfirmationsRepository.GetAllConfirmedEmails(userModel.Id)
	responseModel := email_confirmations.EmailConfirmationsListResponse{
		emailsList,
	}
	resp, _ := json2.Marshal(responseModel)

	return c.String(http.StatusOK, string(resp))
}

func (rec RealEmailsCntrl) Update(c echo.Context) error {
	currentUser := http2.CurrentUser(c)
	userModel := rec.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)
	request := real_emails.ReplaceRealEmailRequest{}

	reqBody, err1 := io.ReadAll(c.Request().Body)
	if err1 != nil {
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

	rec.ReplaceRealEmail.Replace(
		userModel,
		request.OldEmail,
		request.NewEmail,
	)
	resp := common.Success{true}
	json, _ := json2.Marshal(resp)

	return c.String(http.StatusOK, string(json))
}
