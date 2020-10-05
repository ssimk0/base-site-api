package config

import (
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/spf13/viper"
)

func loadCompressionConfiguration() (enabled bool, config compress.Config, err error) {
	// Set a new configuration provider
	provider := viper.New()

	// Set configuration provider settings
	provider.SetConfigName("compress")
	provider.AddConfigPath("./config")

	// Set default configurations
	setDefaultCompressionConfiguration(provider)

	// Read configuration file
	err = provider.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// Config file not found; ignore error since we have default configurations

	} else if err != nil {
		// Config file was found but another error was produced
		return provider.GetBool("Enabled"), config, err
	}

	// Unmarshal the configuration file into logger.Config
	err = provider.Unmarshal(&config)

	// Return the configuration (and error if occurred)
	return provider.GetBool("Enabled"), config, err
}

// Set default configuration for the Logger Middleware
func setDefaultCompressionConfiguration(provider *viper.Viper) {
	provider.SetDefault("Enabled", true)
	provider.SetDefault("Level", 0)
}
