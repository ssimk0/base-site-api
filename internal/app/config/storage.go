package config

import (
	"github.com/spf13/viper"
)

type StorageConfiguration struct {
	Type      string
	Region    string
	AccessKey string
	SecretKey string
	Bucket    string
	Endpoint  string
}

func loadStorageConfiguration() (StorageConfiguration, error) {
	// Set a new configuration provider
	provider := viper.New()

	// Set configuration provider settings
	provider.SetConfigName("storage")
	provider.AddConfigPath("./config")

	// Create a new fiber.Settings variable
	var config StorageConfiguration

	// Set default configurations
	setDefaultStorageConfiguration(provider)

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
func setDefaultStorageConfiguration(provider *viper.Viper) {
	provider.SetDefault("Type", "local")
	provider.SetDefault("Endpoint", "")
	provider.SetDefault("Region", "")
	provider.SetDefault("AccessKey", "")
	provider.SetDefault("SecretKey", "")
	provider.SetDefault("Bucket", "")
}
