package config

import
	"os"

type EnvConfig struct {
	ConfigPath  string
	HttpPort    string
	Environment string
}

func ReadEnvironmentVariables() *EnvConfig {
	config := EnvConfig{
		HttpPort: os.Getenv("PORT"),
		ConfigPath: os.Getenv("CONFIG_PATH"),
		Environment: os.Getenv("ENV"),
	}

	if config.ConfigPath == "" {
		config.ConfigPath = "etc/"
	}

	if config.Environment == "" {
		config.ConfigPath = "live"
	}

	return &config
}
