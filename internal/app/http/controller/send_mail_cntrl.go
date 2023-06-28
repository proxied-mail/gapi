package controller

import (
	"encoding/json"
	"fmt"
	http2 "github.com/abrouter/gapi/internal/app/http"
	"github.com/abrouter/gapi/internal/app/http/request/send_mail"
	"github.com/abrouter/gapi/internal/app/http/response/common"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/pkg/mail_delivery"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	validator2 "gopkg.in/go-playground/validator.v9"
	"io"
	"net/http"
	"strings"
)

type SendMailCntrl struct {
	fx.In
	repository.UserRepository
	repository.CustomDomainsRepository
}

func (smc SendMailCntrl) Create(c echo.Context) error {

	var req send_mail.SendMailRequest
	reqBody, err1 := io.ReadAll(c.Request().Body)
	if err1 != nil {
		resp, _ := json.Marshal(ErrorResponse{
			Message: "Invalid json",
			Status:  false,
		})
		return c.String(http.StatusUnprocessableEntity, string(resp))
	}

	err2 := json.Unmarshal(reqBody, &req)
	if err2 != nil {
		resp, _ := json.Marshal(ErrorResponse{
			Message: "Invalid json",
			Status:  false,
		})
		return c.String(http.StatusUnprocessableEntity, string(resp))
	}

	valid := validator2.New().Struct(req)

	if valid != nil {
		resp, _ := json.Marshal(ErrorResponse{
			Message: valid.Error(),
			Status:  false,
		})
		return c.String(http.StatusUnprocessableEntity, string(resp))
	}

	domain := strings.Split(req.Auth.Username, "@")[1]
	currentUser := http2.CurrentUser(c)
	userModel := smc.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)
	smc.CustomDomainsRepository.UserHasDomain(userModel.Id, domain)

	sendMailResult := mail_delivery.SendMail(req.Auth, req.Mail)
	if sendMailResult != nil {
		fmt.Println(sendMailResult)
		return c.String(http.StatusOK, "Error")
	}

	return c.String(http.StatusOK, common.GetSuccess())
}
