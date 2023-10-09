package models

import (
	"database/sql"
	"time"
)

const RELATED_TO_TYPE_PROXY_BINDING = 1

type Passwords struct {
	ID            int            `json:"id"`
	UserId        int            `json:"user_id"`
	RelatedToType int            `json:"related_to_type"`
	RelatedToId   int            `json:"related_to_id"`
	Login         sql.NullString `json:"login"`
	Website       sql.NullString `json:"website"`
	Password      string         `json:"password"`
	Title         sql.NullString `json:"title"`
	Note          sql.NullString `json:"note"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}
