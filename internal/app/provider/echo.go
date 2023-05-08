package provider

import "github.com/labstack/echo/v4"

func EchoProvider() *echo.Echo {
	return echo.New()
}
