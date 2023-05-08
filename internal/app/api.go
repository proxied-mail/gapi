package app

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ConfigureApiRoutes(e *echo.Echo) {
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
