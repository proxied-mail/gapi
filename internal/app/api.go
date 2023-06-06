package app

import (
	"github.com/abrouter/gapi/internal/app/controller"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ConfigureApiRoutes(e *echo.Echo) {
	e.GET("/gapi/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/gapi/status", controller.Status)
}
