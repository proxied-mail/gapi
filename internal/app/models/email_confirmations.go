package models

import "time"

type EmailConfirmations struct {
	ID                       int    `json:"id"`
	UserId                   int    `json:"user_id"`
	Email                    string `json:"email"`
	RawEmail                 string `json:"raw_email"`
	Code                     string
	Type                     int `json:"type"`
	Confirmed                int
	ShownConfirmationRequest bool `json:"shown_confirmation_request"`
	Attempts                 int
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}
