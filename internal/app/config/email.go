package config

import (
	"github.com/spf13/viper"
)

type EmailConfiguration struct {
	SmtpHost   string
	Port       string
	Username   string
	Password   string
	SecretKey  string
	Encryption string
}

func loadEmailConfiguration() (EmailConfiguration, error) {
	// Set a new configuration provider
	provider := viper.New()

	// Set configuration provider settings
	provider.SetConfigName("email")
	provider.AddConfigPath("./config")

	// Create a new fiber.Settings variable
	var config EmailConfiguration

	// Set default configurations
	setDefaultEmailConfiguration(provider)

	// Read configuration storage
	err := provider.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// Config storage not found; ignore error since we have default configurations
	} else if err != nil {
		// Config storage was found but another error was produced
		return config, err

	}

	// Unmarshal the configuration storage into the ApplicationConfiguration struct
	err = provider.Unmarshal(&config)

	if err != nil {
		return config, err
	}

	// Return the configuration (and error if occurred)
	return config, err
}

// Set default configuration for the application
func setDefaultEmailConfiguration(provider *viper.Viper) {
	provider.SetDefault("SmtpHost", "")
	provider.SetDefault("Port", "")
	provider.SetDefault("Username", "")
	provider.SetDefault("Password", "")
	provider.SetDefault("SecretKey", "")
	provider.SetDefault("Encryption", "")
}
