package controller

import (
	json2 "encoding/json"
	http2 "github.com/abrouter/gapi/internal/app/http"
	"github.com/abrouter/gapi/internal/app/http/response/email_confirmations"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RealEmailsCntrl struct {
	UserRepository               repository.UserRepository
	EmailConfirmationsRepository repository.EmailConfirmationsRepository
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
