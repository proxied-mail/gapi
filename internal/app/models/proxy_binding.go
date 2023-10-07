package models

import "time"

type ProxyBinding struct {
	ID             int       `json:"id"`
	UserId         int       `json:"user_id"`
	ReverseFor     int       `json:"reverse_for"`
	ProxyAddress   string    `json:"proxy_address"`
	DeliveryMethod int       `json:"delivery_method"`
	ReceivedEmails int       `json:"received_emails"`
	AttrsJson      string    `json:"attrs_json"`
	Type           int       `json:"type"`
	Description    string    `json:"description"`
	CallbackUrl    string    `json:"callback_url"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}
