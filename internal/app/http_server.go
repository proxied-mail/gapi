package app

import "github.com/labstack/echo/v4"

func StartHttpServer(e *echo.Echo) {
	e.Start(":9000")
}
