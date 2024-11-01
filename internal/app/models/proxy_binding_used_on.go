package models

import (
	"time"
)

type ProxyBindingUsedOn struct {
	Id             int       `json:"id"`
	UserId         int       `json:"user_id"`
	ProxyBindingId int       `json:"proxy_binding_id"`
	JsonList       string    `json:"json_list"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (ProxyBindingUsedOn) TableName() string {
	return "proxy_binding_used_on"
}
