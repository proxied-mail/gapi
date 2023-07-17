package models

import (
	"database/sql"
	"time"
)

const DomainStatusNew = 1
const DomainStatusOwnershipVerified = 2
const DomainStatusMxSet = 3
const DomainStatusSpfSet = 4
const DomainStatusDkimSet = 5

type CustomDomain struct {
	ID           int            `json:"id"`
	UserId       int            `json:"user_id"`
	Domain       string         `json:"domain"`
	DisplayName  string         `json:"display_name"`
	Status       int            `json:"status"`
	IsShared     bool           `json:"is_shared"`
	IsPremium    bool           `json:"IsPremium"`
	SmtpPassword sql.NullString `json:"SmtpPassword"`
	DkimKey      string         `json:"DkimKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}
