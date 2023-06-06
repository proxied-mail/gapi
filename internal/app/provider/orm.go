package provider

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OrmProvider(db *sql.DB) *gorm.DB {
	orm, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))
	if err != nil {
		panic(err)
	}

	return orm
}
