package mail_delivery

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	easydkim "github.com/abrouter/gapi/pkg/mail_delivery/dkim"
	"gopkg.in/gomail.v2"
	"io"
	"log"
	"net/smtp"
	"strconv"
	"strings"
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

	var err error
	var buffer bytes.Buffer
	_, err = m.WriteTo(&buffer)
	if err != nil {
		log.Fatal(err)
	}

	separated := strings.Split(sendMailCommand.From, "@")
	domain := ""
	if len(separated) > 0 {
		domain = strings.Replace(separated[1], ">", "", 999)
	}

	fmt.Println("domain111:" + domain + "'")

	privateKeyPath := "/app/config/dkim/key.private"
	message := buffer.Bytes()
	if domain != "" {
		var signedMessage []byte
		signedMessage, err = easydkim.Sign(
			buffer.Bytes(),
			privateKeyPath,
			"dkim",
			domain,
		)
		if err != nil {
			fmt.Println("DKIM signing error" + err.Error())
			fmt.Println("DKIM wasn't signed")
		} else {
			message = signedMessage
			fmt.Println(string(message))
		}
	}

	auth := smtp.PlainAuth("", authData.Username, authData.Password, authData.Host)
	err = smtp.SendMail(
		authData.Host+":"+strconv.Itoa(authData.Port),
		auth,
		sendMailCommand.From,
		[]string{
			sendMailCommand.To,
		},
		message,
	)

	if err != nil {
		return err
	}

	return nil
}
