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
	e.GET("/gapi/domains", api.DomainsController.List, middlewares.AuthMiddleware)
	e.GET("/gapi/verified-emails-list", api.RealEmailsCntrl.Get, middlewares.AuthMiddleware)
}
