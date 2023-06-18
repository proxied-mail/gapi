package boot

import "flag"

type AppFlags struct {
	ConfigFilePath string
}

var Flags AppFlags

func ParseFlags() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "../../config/.env-local", "path to config file")
	flag.Parse()
	Flags = AppFlags{
		ConfigFilePath: configFilePath,
	}
}

func GetConfigFilePath() string {
	return Flags.ConfigFilePath
}
