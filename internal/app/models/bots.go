package models

type Bots struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Uid    string `json:"uid"`
	Name   string `json:"name"`
}
