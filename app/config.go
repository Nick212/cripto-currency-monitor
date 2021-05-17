package app

import (
	"github.com/spf13/viper"
)

// Config contains the configurations for application.
type Config struct {
	// Application version.
	Version string

	// Application Culture.
	Culture string

	// HOST CRM
	HOST_API string

	// IndexName
	IndexName string
}

// InitConfig load the configuration for application.
func InitConfig() (*Config, error) {
	config := &Config{
		Version:   viper.GetString("VERSION"),
		Culture:   viper.GetString("CULTURE"),
		HOST_API:  viper.GetString("HOST_API"),
		IndexName: viper.GetString("INDEX_NAME"),
	}

	return config, nil
}
