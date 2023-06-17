package provider

import (
	"database/sql"
	"github.com/abrouter/gapi/internal/app"
)
import "fmt"
import _ "github.com/go-sql-driver/mysql"

type MysqlRwConnectionProvider struct {
	Connection *sql.DB
}

func (m MysqlRwConnectionProvider) Connect() *sql.DB {
	fmt.Println("Starting connection to MySQL")

	if m.Connection != nil {
		return m.Connection
	}

	connection, err := sql.Open("mysql", app.GetMysqlConnectionString())
	if err != nil {
		panic(err.Error())
	}
	m.Connection = connection
	fmt.Println("MySQL successfully connected")

	return m.Connection
}
