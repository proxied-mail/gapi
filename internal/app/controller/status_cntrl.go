package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

type StatusEntity struct {
	Status bool
}

func Status(c echo.Context) error {
	resp, _ := json.Marshal(StatusEntity{Status: true})

	return c.String(http.StatusOK, string(resp))
}
