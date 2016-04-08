package config

import (
	//"os"
	"gopkg.in/alecthomas/kingpin.v2"
	//"fmt"
)

type CliConfig struct {
	ConfigPath  *string
	HttpPort    *int
	Environment *string
}

var (
	env     = kingpin.Flag("env", "Current environment: dev or live").Default("live").HintOptions("dev", "live").Short('e').String()
	port    = kingpin.Arg("port", "Port to run application").Int()
	config  = kingpin.Arg("config", "Config path, etc/").String()
)


func ReadCliArgs() *CliConfig {

	kingpin.Version("0.0.1")
	kingpin.Parse()

	config := CliConfig{
		HttpPort: port,
		ConfigPath: config,
		Environment: env,
	}

	return &config
}
