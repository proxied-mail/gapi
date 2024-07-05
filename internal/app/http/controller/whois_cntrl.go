package controller

import (
	"encoding/json"
	"github.com/haccer/available"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type WhoisCntrl struct {
	fx.In
}

type WhoisResponse struct {
	IsDomainRegistered bool `json:"is_domain_registered"`
	HasError           bool `json:"has_error"`
}

func (cntrl WhoisCntrl) Whois(c echo.Context) error {

	domain := c.QueryParam("domain")
	isDomainRegisterred := available.Domain(domain)

	resp, _ := json.Marshal(WhoisResponse{
		IsDomainRegistered: isDomainRegisterred,
		HasError:           false,
	})

	return c.String(200, string(resp))
}
