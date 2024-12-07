package received_emalis

import (
	"encoding/json"
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/pkg/entityId"
	"go.uber.org/fx"
)

type ReceivedEmailDTO struct {
	Id               string        `json:"id"`
	UseThisIdToReply string        `json:"useThisIdToReply"`
	Sender           string        `json:"sender"`
	Recipient        string        `json:"recipient"`
	Subject          string        `json:"subject"`
	BodyHtml         string        `json:"bodyHtml"`
	BodyPlain        string        `json:"bodyPlain"`
	Attachments      []Attachments `json:"attachments"`
}

type Attachments struct {
	Filename string `json:"filename"`
	Url      string `json:"url"`
}

type ReceivedEmailParser struct {
	fx.In
	entityId.Encoder
}

func (r ReceivedEmailParser) Parse(m models.ReceivedEmails) (ReceivedEmailDTO, error) {
	var payload ReceivedEmailDTO
	err := json.Unmarshal([]byte(m.Payload.String), &payload)
	if err != nil {
		return payload, err
	}
	payload.Id = r.Encoder.Encode(m.Id, "received_emails")
	payload.UseThisIdToReply = payload.Id

	return payload, nil
}
