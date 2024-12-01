package models

import (
	"database/sql"
	"time"
)

type ReceivedEmails struct {
	Id             int            `json:"id"`
	Payload        sql.NullString `json:"payload"`
	IsProcessed    int            `json:"is_processed"`
	RecipientEmail string         `json:"recipient_email"`
	SenderEmail    string         `json:"sender_email"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}
