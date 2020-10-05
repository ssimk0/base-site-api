package config

import (
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/spf13/viper"
)

func loadRecoverConfiguration() (enabled bool, config recover.Config, err error) {
	// Set a new configuration provider
	provider := viper.New()

	// Set configuration provider settings
	provider.SetConfigName("recover")
	provider.AddConfigPath("./config")

	// Set default configurations
	setDefaultRecoverConfiguration(provider)

	// Read configuration file
	err = provider.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error since we have default configurations
		} else {
			// Config file was found but another error was produced
			return provider.GetBool("Enabled"), config, err
		}
	}

	// Unmarshal the configuration file into recover.Config
	err = provider.Unmarshal(&config)

	// Return the configuration (and error if occurred)
	return provider.GetBool("Enabled"), config, err
}

// Set default configuration for the Recover Middleware
func setDefaultRecoverConfiguration(provider *viper.Viper) {
	provider.SetDefault("Enabled", true)
}
