package provider

import "os"

func SetEnvDefaults() {
	if os.Getenv("DB_CONNECTION") == "" {
		os.Setenv("DB_CONNECTION", "root:example@tcp(127.0.0.1:33072)/pm?parseTime=true")
	}

	if os.Getenv("PAPI_HOST") == "" {
		os.Setenv("PAPI_HOST", "http://localhost:904")
	}

}
