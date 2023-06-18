package env

import (
	"github.com/abrouter/gapi/internal/app/boot"
	"github.com/hashicorp/go-envparse"
	"os"
	"strings"
)

var env map[string]string

func ReadEnv() map[string]string {
	file, err := os.ReadFile(boot.GetConfigFilePath())
	if err != nil {
		panic("Error reading .env file by path" + boot.GetConfigFilePath())
	}
	fileContent := string(file)

	var err2 error
	env, err2 = envparse.Parse(strings.NewReader(fileContent))
	if err2 != nil {
		panic(err2)
	}

	return env
}

func GetMysqlConnectionString() string {
	return env["DB_CONNECTION"]
}

func GetPapiHost() string {
	return env["PAPI_HOST"]
}
