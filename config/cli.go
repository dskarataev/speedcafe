package config

import (
	"flag"
)

type CliConfig struct {
	ConfigPath  string
	HttpPort    string
	Environment string
}

func ReadCliArgs() *CliConfig {

	config := CliConfig{
		HttpPort: *flag.String("port", "", "Http port"),
		ConfigPath: *flag.String("config", "", "Path to config file"),
		Environment: *flag.String("env", "", "Environment"),
	}

	flag.Parse();

	return &config
}
