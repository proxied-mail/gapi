package models

import "time"

type RealAddress struct {
	ID                   int       `json:"id"`
	UserId               int       `json:"user_id"`
	ProxyBindingId       int       `json:"proxy_binding_id"`
	RealAddress          string    `json:"real_address"`
	IsEnabled            int       `json:"is_enabled"`
	IsVerificationNeeded int       `json:"is_verification_needed"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	DeletedAt            time.Time `json:"deleted_at"`
}

func (RealAddress) TableName() string {
	return "real_addresses_groups"
}
