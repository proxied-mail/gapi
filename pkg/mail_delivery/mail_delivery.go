package mail_delivery

import (
	"gopkg.in/gomail.v2"
)

type SendMailAuthData struct {
	Host     string `json:"host" validate:"required"`
	Port     int    `json:"port" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SendMailCommand struct {
	From    string `json:"from" validate:"required"`
	To      string `json:"to" validate:"required"`
	Subject string `json:"subject" validate:"required"`
	Type    string `json:"type" validate:"required"`
	Body    string `json:"body" validate:"required"`
	ReplyTo string `json:"reply_to"`
}

func SendMail(authData SendMailAuthData, sendMailCommand SendMailCommand) error {
	m := gomail.NewMessage()
	m.SetHeader("From", sendMailCommand.From)
	m.SetHeader("To", sendMailCommand.To)
	m.SetHeader("Subject", sendMailCommand.Subject)
	if sendMailCommand.ReplyTo != "" {
		m.SetHeader("Reply-To", sendMailCommand.ReplyTo)
	}
	m.SetBody(sendMailCommand.Type, sendMailCommand.Body)

	d := gomail.NewDialer(authData.Host, authData.Port, authData.Username, authData.Password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
