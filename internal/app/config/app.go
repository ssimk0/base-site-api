package config

import (
	"github.com/spf13/viper"
	"os"
	"path"
)

type ApplicationConfiguration struct {
	Listen         string
	SigningKeyPath string
	SigningKey     []byte
	TemplatePath   string
	SentryDNS      string
	AppURL         string
	Env            string
	LogToFile      bool
	Debug          bool
}

func loadApplicationConfiguration() (ApplicationConfiguration, error) {
	// Set a new configuration provider
	provider := viper.New()

	// Set configuration provider settings
	provider.SetConfigName("app")
	provider.AddConfigPath("./config")

	// Create a new fiber.Settings variable
	var config ApplicationConfiguration

	// Set default configurations
	setDefaultApplicationConfiguration(provider)

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

	dir, err := os.Getwd()

	if err != nil {
		return config, err
	}

	config.SigningKey, err = os.ReadFile(path.Join(dir, config.SigningKeyPath))
	if err != nil {
		return config, err
	}

	// Return the configuration (and error if occurred)
	return config, err
}

// Set default configuration for the application
func setDefaultApplicationConfiguration(provider *viper.Viper) {
	provider.SetDefault("Listen", "8080")
	provider.SetDefault("AppURL", "http://localhost:3000/")
	provider.SetDefault("TemplatePath", "templates")
	provider.SetDefault("Debug", false)
	provider.SetDefault("Env", "production")
	provider.SetDefault("LogToFile", false)
	provider.SetDefault("SigningKeyPath", "./jwtRS256.key")
}
