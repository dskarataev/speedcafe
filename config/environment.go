package config

import
	"os"

type EnvConfig struct {
	ConfigPath  string
	HttpPort    string
	Environment string
}

func ReadEnvironmentVariables() *EnvConfig {
	var config EnvConfig
	config.HttpPort = os.Getenv("PORT")
	config.ConfigPath = os.Getenv("CONFIG_PATH")
	config.Environment = os.Getenv("ENV")

	if config.ConfigPath == "" {
		config.ConfigPath = "etc/"
	}

	if config.Environment == "" {
		config.ConfigPath = "live"
	}

	return &config
}
