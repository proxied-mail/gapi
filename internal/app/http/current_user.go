package http

import (
	"github.com/abrouter/gapi/pkg/papi"
	"github.com/labstack/echo/v4"
)

func CurrentUser(c echo.Context) papi.PapiUserStruct {
	auth := c.Request().Header.Get("Authorization")
	c.Response().Header().Add("content-type", "application/json")
	me, _ := papi.MeCached(auth)
	return me
}
