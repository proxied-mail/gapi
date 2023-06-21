package models

import "time"

type EmailConfirmations struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Email     string `json:"email"`
	Code      string
	Confirmed int
	Attempts  int
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
