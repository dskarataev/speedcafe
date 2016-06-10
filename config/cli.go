package config

import (
	"flag"
	"fmt"
)

type CliConfig struct {
	ConfigPath  string
	HttpPort    string
	Environment string
}

func ReadCliArgs() *CliConfig {

	Environment := flag.String("env", "", "Environment")
	ConfigPath := flag.String("config", "", "Path to config file")
	HttpPort := flag.String("port", "8888", "Http port")

	flag.Parse();

	config := CliConfig{
		HttpPort: *HttpPort,
		ConfigPath: *ConfigPath,
		Environment: *Environment,
	}

	fmt.Println("Cli Variables: ")
	fmt.Print("     HttpPort: " + config.HttpPort)
	fmt.Print(", ConfigPath: " + config.ConfigPath)
	fmt.Print(", Environment: " + config.Environment)
	fmt.Println("")

	return &config
}
