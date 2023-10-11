package models

import (
	"time"
)

type ProxyBindingUsedOn struct {
	Id             int       `json:"id"`
	UserId         int       `json:"user_id"`
	ProxyBindingId int       `json:"proxy_binding_id"`
	JsonList       int       `json:"json_list"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
