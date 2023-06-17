package env

import "os"

func GetMysqlConnectionString() string {
	return os.Getenv("DB_CONNECTION")
}

func GetPapiHost() string {
	return os.Getenv("PAPI_HOST")
}
