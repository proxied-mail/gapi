package mail_delivery

import (
	"bytes"
	"crypto/md5"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	easydkim "github.com/abrouter/gapi/pkg/mail_delivery/dkim"
	"gopkg.in/gomail.v2"
	"io"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"time"
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

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func addLineBreaks(input string, maxLineLength int) string {
	var result strings.Builder
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		words := strings.Fields(line)
		lineLength := 0

		for _, word := range words {
			wordLength := len(word)
			if lineLength+wordLength > maxLineLength {
				result.WriteString("\n")
				lineLength = 0
			}
			if lineLength > 0 {
				result.WriteString(" ")
				lineLength++
			}
			result.WriteString(word)
			lineLength += wordLength
		}
		result.WriteString("\n") // Preserve original line breaks
	}

	return result.String()
}

func messageId() string {
	min := 1000
	max := 9999
	rand := rand.Intn(max-min) + min
	val := strconv.FormatInt(time.Now().UnixNano(), 16) + strconv.Itoa(rand)
	val = getMD5Hash(val)
	return "<" + val + "@mx.proxiedmail.com>"
}

func SendMail(authData SendMailAuthData, sendMailCommand SendMailCommand) error {
	m := gomail.NewMessage()

	m.SetHeader("Subject", sendMailCommand.Subject)

	if sendMailCommand.ReplyTo != "" {
		m.SetHeader("Reply-To", sendMailCommand.ReplyTo)
	}
	m.SetHeader("Message-Id", messageId())

	m.SetHeader("Date", time.Now().Format(time.RFC1123Z))

	fmt.Println(sendMailCommand.From)

	from := sendMailCommand.From
	from = strings.Replace(from, "\"", "", 2)
	spaceBeforeLrv := strings.Index(from, "<")
	if spaceBeforeLrv != 0 {
		from = "\"" + from
		from = strings.Replace(from, " <", " \" <", 1)
	}

	m.SetHeader("From", from)
	m.SetHeader("To", sendMailCommand.To)
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
	message = []byte(strings.Replace(string(message), "Mime-Version: 1.0", "MIME-Version: 1.0", 1))

	if domain != "" {
		var signedMessage []byte
		signedMessage, err = easydkim.Sign(
			message,
			privateKeyPath,
			"dkim",
			domain,
		)
		if err != nil {
			fmt.Println("DKIM signing error" + err.Error())
			fmt.Println("DKIM wasn't signed")
		} else {
			message = signedMessage

			//if Is7Bit(message) {
			//	message = convert7bitTo8bit(message)
			//}

			message = []byte(b64.StdEncoding.EncodeToString(message))
			message, _ = b64.StdEncoding.DecodeString(string(message))
			//json encode message into single array to l
			json5, _ := json.Marshal(
				[]string{string(message)},
			)

			fmt.Println(string(json5))
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

func Is7Bit(data []byte) bool {
	for _, b := range data {
		if b > 127 {
			return false
		}
	}
	return true
}

func convert7bitTo8bit(data []byte) []byte {
	// Store the 8-bit ASCII in res.
	var res = make([]byte, len(data)*8/7)

	fmt.Printf("First 100 bytes of Data: % #x\n", data[:100])
	var idx, shift int
	for i, offset := 0, 0; i < len(res)-1; i, offset = i+1, offset+7 {
		idx, shift = offset/8, offset%8

		lhs := data[idx] & (255 >> shift)
		if shift == 0 {
			lhs >>= 1
		} else if shift > 1 {
			lhs <<= shift - 1
		}
		rhs := (data[idx+1] & (255 << (9 - shift))) >> (9 - shift)
		res[i] = (lhs | rhs) & 127
	}
	fmt.Printf("First 100 bytes encoded: % #x\nFirst 100 chars: %s", res[:100], res[:100])
	return res
}
