package app

import (
	"github.com/spf13/viper"
)

// Config contains the configurations for application.
type Config struct {
	// Application version.
	Version string

	// HOST CRM
	HOST_API []string
}

// InitConfig load the configuration for application.
func InitConfig() (*Config, error) {
	config := &Config{
		Version:                  viper.GetString("VERSION"),
		HOST_API:             viper.GetStringSlice("HOST_API"),
	}

	return config, nil
}
