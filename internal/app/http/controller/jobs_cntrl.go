package controller

import (
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"net/http"
)

type JobsController struct {
	fx.In
	jbsRep repository.JobsRepository
}

func (jbCntrl JobsController) status(c echo.Context) error {
	count := jbCntrl.jbsRep.Count()
	status := http.StatusOK
	if count > 20 {
		status = http.StatusInternalServerError
	}

	return c.String(status, "nil")
}
