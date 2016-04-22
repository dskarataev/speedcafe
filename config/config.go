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

func (this *Config) Init() error {

	envVariables := ReadEnvironmentVariables()
	cliArgs := ReadCliArgs()

	cfg := ini.Empty()
	// it makes reading 50-70% faster, but we should not write to config file in that case
	cfg.BlockMode = false

	fmt.Println("Result: ")
	// Calculate ConfigPath
	configPath := choose(envVariables.ConfigPath, cliArgs.ConfigPath);
	err := cfg.Append(configPath + "app.ini")
	fmt.Println("   Config: " + configPath + "app.ini")

	if err != nil {
		return err
	}

	env := choose(cliArgs.Environment, envVariables.Environment);
	fmt.Println("   Environment: " + env)

	if env != "live" {
		err := cfg.Append(configPath + env + ".ini")
		fmt.Println("   Additional config: " + configPath + env + ".ini")
		if err != nil {
			return err
		}
	}

	cfg.Section("").NewKey("http_port", choose(envVariables.HttpPort, cliArgs.HttpPort));

	err = cfg.MapTo(this)
	if err != nil {
		return err
	}

	fmt.Println("Configurations:")
	fmt.Println("   HTTP port: " + this.HttpPort)
	fmt.Println("   Database host: " + this.Database.Host + ", name: " + this.Database.Name + ", user: " + this.Database.User + ", password: " + this.Database.Password)
	fmt.Println("   Foursquare id: " + this.FoursquareClient.ID + ", key: " + this.FoursquareClient.SecretKey)
	fmt.Println("")
	return nil
}

func choose(first string, second string) string {
	if first != "" {
		return first
	} else {
		return second
	}
}
