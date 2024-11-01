package controller

import (
	json2 "encoding/json"
	http2 "github.com/abrouter/gapi/internal/app/http"
	"github.com/abrouter/gapi/internal/app/http/request/real_emails"
	"github.com/abrouter/gapi/internal/app/http/response/common"
	"github.com/abrouter/gapi/internal/app/http/response/email_confirmations"
	"github.com/abrouter/gapi/internal/app/http/response/real_emails_rsp"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/internal/app/services/real_emails_srv"
	"github.com/abrouter/gapi/pkg/entityId"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	validator2 "gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type RealEmailsCntrl struct {
	fx.In
	UserRepository               repository.UserRepository
	EmailConfirmationsRepository repository.EmailConfirmationsRepository
	RealEmailsRepository         repository.RealEmailsRepository
	ReplaceRealEmail             real_emails_srv.ReplaceRealEmail
	Encoder                      entityId.Encoder
	Db                           *gorm.DB
}

func (rec RealEmailsCntrl) GetAll(c echo.Context) error {
	currentUser := http2.CurrentUser(c)
	userModel := rec.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)

	emailsList := rec.RealEmailsRepository.GetAllUniqueByUser(userModel.Id)
	responseModel := real_emails_rsp.MapResponse(emailsList)

	resp, _ := json2.Marshal(responseModel)

	return c.String(http.StatusOK, string(resp))
}

func (rec RealEmailsCntrl) GetVerified(c echo.Context) error {
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

func (rec RealEmailsCntrl) MarkAsVerificationRequestShown(c echo.Context) error {
	currentUser := http2.CurrentUser(c)
	userModel := rec.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)

	request := real_emails.MarkAsVerReqShownRequest{}
	reqBody, err := io.ReadAll(c.Request().Body)
	if err != nil {
		resp, _ := json2.Marshal(ErrorResponse{
			Message: "Invalid json",
			Status:  false,
		})
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}
	json2.Unmarshal(reqBody, &request)
	id, err2 := rec.Encoder.Decode(request.Id, "email_confirmations")
	if err2 != nil {
		resp := ErrorResponse{
			Message: "Error on decoding entity",
			Status:  false,
		}
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	confirmation := rec.EmailConfirmationsRepository.GetByIdAndUserId(int(id), userModel.Id)
	if confirmation.ID < 1 {
		resp := ErrorResponse{
			Message: "Cant find confirmation",
			Status:  false,
		}
		return c.JSON(http.StatusNotFound, resp)
	}
	confirmation.ShownConfirmationRequest = true
	rec.Db.Save(confirmation)

	resp := email_confirmations.MarkAsShownResponse{
		Status: true,
		Email:  confirmation.RawEmail,
	}

	return c.JSON(http.StatusOK, resp)
}

func (rec RealEmailsCntrl) CheckEmailConfirmation(c echo.Context) error {
	currentUser := http2.CurrentUser(c)
	userModel := rec.UserRepository.GetUserByEmail(currentUser.Data.Attributes.Username)
	firstUnconfirmed := rec.EmailConfirmationsRepository.FirstUnconfirmedNotShown(userModel.Id)
	hasConfirmedEmails := rec.EmailConfirmationsRepository.HasConfirmedEmails(userModel.Id)

	if firstUnconfirmed.ID < 1 || hasConfirmedEmails {
		response := email_confirmations.FirstUnconfirmedResponse{
			HasUnconfirmedNotShown: false,
			Id:                     "",
			ContinueChecking:       false,
		}
		return c.JSON(http.StatusOK, response)
	}

	id := rec.Encoder.Encode(firstUnconfirmed.ID, "email_confirmations")
	response := email_confirmations.FirstUnconfirmedResponse{
		HasUnconfirmedNotShown: firstUnconfirmed.ID > 0,
		Id:                     id,
		ContinueChecking:       true,
	}

	return c.JSON(http.StatusOK, response)
}
