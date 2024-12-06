package app

import (
	"github.com/abrouter/gapi/internal/app/http/controller"
	"github.com/abrouter/gapi/internal/app/http/middlewares"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"net/http"
)

type api struct {
	fx.In
	controller.DomainsController
	controller.RealEmailsCntrl
	controller.SendMailCntrl
	controller.PasswordsCntrl
	controller.UsedOnCntrl
	controller.JobsController
	controller.UiTestController
	controller.WhoisCntrl
	controller.SettingsController
	controller.BotController
	controller.ConversationsController
	controller.AssignBotController
	controller.UpdateBotController
	controller.PbBotGetController
}

func ConfigureApiRoutes(
	e *echo.Echo,
	api api,
) {
	e.GET("/gapi/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/gapi/status", controller.Status)
	e.POST("/gapi/domains", api.DomainsController.Create, middlewares.AuthMiddleware)
	e.POST("/gapi/send-mail", api.SendMailCntrl.Create, middlewares.AuthMiddleware)
	e.POST("/gapi/real-emails/replace", api.RealEmailsCntrl.Update, middlewares.AuthMiddleware)
	e.GET("/gapi/available-domains", api.DomainsController.List, middlewares.AuthMiddleware)
	e.GET("/gapi/custom-domains", api.DomainsController.ListCustom, middlewares.AuthMiddleware) //custom domains
	e.GET("/gapi/verified-emails-list", api.RealEmailsCntrl.GetVerified, middlewares.AuthMiddleware)
	e.POST(
		"/gapi/email-confirmations/mark-as-req-page-shown",
		api.RealEmailsCntrl.MarkAsVerificationRequestShown,
		middlewares.AuthMiddleware,
	)
	e.GET(
		"/gapi/email-confirmations/check",
		api.RealEmailsCntrl.CheckEmailConfirmation,
		middlewares.AuthMiddleware,
	)

	e.GET("/gapi/real-emails", api.RealEmailsCntrl.GetAll, middlewares.AuthMiddleware)

	e.GET("/gapi/settings", api.SettingsController.GetAll, middlewares.AuthMiddleware)
	e.PATCH("/gapi/settings/update", api.SettingsController.Update, middlewares.AuthMiddleware)

	e.PATCH("/gapi/passwords/proxy-binding", api.PasswordsCntrl.Update, middlewares.AuthMiddleware)
	e.GET("/gapi/passwords", api.PasswordsCntrl.List, middlewares.AuthMiddleware)

	e.PATCH("/gapi/used-on", api.UsedOnCntrl.Change, middlewares.AuthMiddleware)
	e.GET("/gapi/used-on", api.UsedOnCntrl.List, middlewares.AuthMiddleware)

	e.GET("/gapi/whois", api.WhoisCntrl.Whois)

	e.GET("/gapi/jobs-status", api.JobsController.Status)
	e.GET("/gapi/basic-ui-test", api.UiTestController.Basic)

	/**`
	bots_req
	*/
	e.POST(
		"/internal/proxy-binding-bots-req/notify/received-email",
		api.BotController.ReceivedEmailNotify,
	)
	e.POST(
		"/gapi/proxy-binding-bots/assign",
		api.AssignBotController.AssignBot,
	)
	e.PATCH(
		"/gapi/proxy-binding-bots/bot",
		api.UpdateBotController.UpdateBot,
	)
	e.GET(
		"/gapi/bot/conversations",
		api.ConversationsController.GetMessages,
	)

	e.GET(
		"/gapi/proxy-binding-bots/get",
		api.PbBotGetController.Get,
	)

}
