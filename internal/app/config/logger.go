package config

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/spf13/viper"
)

func loadLoggerConfiguration() (enabled bool, config logger.Config, err error) {
	// Set a new configuration provider
	provider := viper.New()

	// Set configuration provider settings
	provider.SetConfigName("logger")
	provider.AddConfigPath("./config")

	// Set default configurations
	setDefaultLoggerConfiguration(provider)

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
func setDefaultLoggerConfiguration(provider *viper.Viper) {
	provider.SetDefault("LoggerEnabled", true)
	provider.SetDefault("Filter", nil)
	provider.SetDefault("Format", "${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n")
	provider.SetDefault("TimeFormat", "15:04:05")
	provider.SetDefault("Output", os.Stderr)
}
