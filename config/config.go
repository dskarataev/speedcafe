package config

import "github.com/go-ini/ini"

type FoursquareClient struct {
	ID        string `ini:"foursquare_id"`
	SecretKey string `ini:"foursquare_secret"`
}

type Config struct {
	FoursquareClient `ini:"DEFAULT"`
}

func NewConfig() *Config {
	return &Config{}
}

func (this *Config) Init(configPath string) error {
	cfg := ini.Empty()
	// it makes reading 50-70% faster, but we should not write to config file in that case
	cfg.BlockMode = false

	err := cfg.Append(configPath)
	if err != nil {
		return err
	}

	err = cfg.MapTo(this)
	if err != nil {
		return err
	}

	return nil
}
