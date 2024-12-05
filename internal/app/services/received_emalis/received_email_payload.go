package received_emalis

import (
	"encoding/json"
	"github.com/abrouter/gapi/internal/app/models"
	"go.uber.org/fx"
)

type ReceivedEmailDTO struct {
	Sender      string        `json:"sender"`
	Recipient   string        `json:"recipient"`
	Subject     string        `json:"subject"`
	BodyHtml    string        `json:"body-html"`
	BodyPlain   string        `json:"body-plain"`
	Attachments []Attachments `json:"attachments"`
}

type Attachments struct {
	Filename string `json:"filename"`
	Url      string `json:"url"`
}

type ReceivedEmailParser struct {
	fx.In
}

func (r ReceivedEmailParser) Parse(m models.ReceivedEmails) (ReceivedEmailDTO, error) {
	var payload ReceivedEmailDTO
	err := json.Unmarshal([]byte(m.Payload.String), &payload)
	if err != nil {
		return payload, err
	}

	return payload, nil
}
