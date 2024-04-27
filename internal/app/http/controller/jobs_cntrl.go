package controller

import (
	"encoding/json"
	"github.com/abrouter/gapi/internal/app/http/response/jobs"
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
	httpStatus := http.StatusOK
	statusText := "ok"
	if count > 20 {
		httpStatus = http.StatusInternalServerError
		statusText = "fail"
	}
	resp, _ := json.Marshal(jobs.StatusJobsResponse{
		Status: statusText,
		Count:  count,
	})

	return c.String(httpStatus, string(resp))
}
