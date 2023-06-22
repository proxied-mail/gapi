package send_mail

import "github.com/abrouter/gapi/pkg/mail_delivery"

type SendMailRequest struct {
	Auth mail_delivery.SendMailAuthData `json:"auth" validate:"required"`
	Mail mail_delivery.SendMailCommand  `json:"mail" validate:"required"`
}
