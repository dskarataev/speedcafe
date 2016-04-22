package config

import (
	"os"
	"fmt"
)

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
		config.Environment = "live"
	}

	fmt.Println("Environment Variables: ")
	fmt.Print("     HttpPort: " + config.HttpPort)
	fmt.Print(", ConfigPath: " + config.ConfigPath)
	fmt.Print(", Environment: " + config.Environment)
	fmt.Println("")

	return &config
}
