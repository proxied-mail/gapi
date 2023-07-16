package mail_delivery

import (
	b64 "encoding/base64"
	"fmt"
	"gopkg.in/gomail.v2"
	"io"
)

type SendMailAuthData struct {
	Host     string `json:"host" validate:"required"`
	Port     int    `json:"port" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SendMailCommand struct {
	From        string       `json:"from" validate:"required"`
	To          string       `json:"to" validate:"required"`
	Subject     string       `json:"subject" validate:"required"`
	Type        string       `json:"type" validate:"required"`
	Body        string       `json:"body" validate:"required"`
	ReplyTo     string       `json:"reply_to"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Name     string `json:"name"`
	MimeType string `json:"mimeType"`
	Size     int    `json:"size"`
	Content  string `json:"content"`
	Url      string `json:"url"`
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
	for _, attachment := range sendMailCommand.Attachments {
		content, _ := b64.StdEncoding.DecodeString(attachment.Content)
		fmt.Println("decoding base64 bla bla")
		m.Attach(
			attachment.Name,
			gomail.SetCopyFunc(func(w io.Writer) error {
				_, err := w.Write(content)
				return err
			}),
			gomail.SetHeader(map[string][]string{"Content-Type": {attachment.MimeType}}),
		)
	}

	d := gomail.NewDialer(authData.Host, authData.Port, authData.Username, authData.Password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
