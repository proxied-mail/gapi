package middlewares

import (
	"encoding/json"
	intHttp "github.com/abrouter/gapi/internal/app/http"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UnauthenticatedRequest struct {
	Status  bool
	Message string
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		me := intHttp.CurrentUser(c)

		if !me.IsAuthenticated() {
			response, _ := json.Marshal(UnauthenticatedRequest{
				Status:  false,
				Message: "Unauthenticated",
			})
			return c.String(http.StatusUnauthorized, string(response))
		}

		return next(c)
	}
}
