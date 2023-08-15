package real_emails_rsp

import "github.com/abrouter/gapi/internal/app/models"

type RealEmails struct {
	Email string `json:"email"`
}

type RealEmailsResponse struct {
	RealEmails []RealEmails `json:"data"`
}

func MapResponse(models []models.RealAddress) RealEmailsResponse {
	var realEmails []RealEmails

	for _, model := range models {
		realEmails = append(realEmails, RealEmails{
			Email: model.RealAddress,
		})
	}
	rsp := RealEmailsResponse{
		RealEmails: realEmails,
	}

	return rsp
}
