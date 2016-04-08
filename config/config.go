package config

import (
	"github.com/go-ini/ini"
	"fmt"
)

type FoursquareClient struct {
	ID        string `ini:"foursquare_id"`
	SecretKey string `ini:"foursquare_secret"`
}

type Database struct {
	Host      string `ini:"database_host"`
	Name      string `ini:"database_name"`
	User      string `ini:"database_user"`
	Password  string `ini:"database_password"`
}

type Config struct {
	FoursquareClient `ini:"DEFAULT"`
	Database `ini:"DEFAULT"`
	HttpPort string `ini:"http_port"`
}

func NewConfig() *Config {
	return &Config{}
}

func (this *Config) Init(configPath string, env string) error {

	environmentVariables := ReadEnvironmentVariables()
	cliArgs := ReadCliArgs()

	fmt.Print(cliArgs)

	cfg := ini.Empty()
	// it makes reading 50-70% faster, but we should not write to config file in that case
	cfg.BlockMode = false

	err := cfg.Append(environmentVariables.ConfigPath + "app.ini")
	if err != nil {
		return err
	}
	if env != "live" {
		err := cfg.Append(environmentVariables.ConfigPath + environmentVariables.Environment + ".ini")
		if err != nil {
			return err
		}
	}

	if environmentVariables.HttpPort != "" {
		cfg.Section("").NewKey("http_port", environmentVariables.HttpPort);
	}


	err = cfg.MapTo(this)
	if err != nil {
		return err
	}

	return nil
}
