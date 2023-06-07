package models

import "database/sql"

type UserModel struct {
	Id          int
	Username    string
	Password    string
	GoogleId    string
	IsTemporary bool
	CountryCode sql.NullString
	CreatedAt   string
	UpdatedAt   string
}

func (UserModel) TableName() string {
	return "users"
}
