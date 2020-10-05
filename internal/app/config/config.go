package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)

type Config struct {
	Fiber          fiber.Config
	App            ApplicationConfiguration
	Enabled        map[string]bool
	Logger         logger.Config
	Recover        recover.Config
	TemplateEngine func(raw string, bind interface{}) (out string, err error)
	Compression    compress.Config
	CORS           cors.Config
	Helmet         helmet.Config
	PublicPrefix   string
	PublicRoot     string
	Public         fiber.Static
	Database       DatabaseConfiguration
}

func Load() (config Config, err error) {
	config.Enabled = make(map[string]bool)
	// Load the Fiber application configuration
	fiberSettings, err := loadFiberConfiguration()
	if err != nil {
		return config, err
	}
	config.Fiber = fiberSettings

	// Load the application configuration
	appConfig, err := loadApplicationConfiguration()
	if err != nil {
		return config, err
	}
	config.App = appConfig

	// Load the logger middleware configuration
	loggerEnabled, loggerConfig, err := loadLoggerConfiguration()
	if err != nil {
		return config, err
	}
	config.Enabled["logger"] = loggerEnabled
	config.Logger = loggerConfig

	// Load the recover middleware configuration
	recoverEnabled, recoverConfig, err := loadRecoverConfiguration()
	if err != nil {
		return config, err
	}
	config.Enabled["recover"] = recoverEnabled
	config.Recover = recoverConfig

	// Load the compression middleware configuration
	compressionEnabled, compressionConfig, err := loadCompressionConfiguration()
	if err != nil {
		return config, err
	}
	config.Enabled["compression"] = compressionEnabled
	config.Compression = compressionConfig

	// Load the CORS middleware configuration
	corsEnabled, corsConfig, err := loadCORSConfiguration()
	if err != nil {
		return config, err
	}
	config.Enabled["cors"] = corsEnabled
	config.CORS = corsConfig

	// Load the Helmet middleware configuration
	helmetEnabled, helmetConfig, err := loadHelmetConfiguration()
	if err != nil {
		return config, err
	}
	config.Enabled["helmet"] = helmetEnabled
	config.Helmet = helmetConfig

	// Load the database configuration
	databaseConfig, err := loadDatabaseConfiguration()
	if err != nil {
		return config, err
	}

	config.Database = databaseConfig

	// Return the configuration
	return config, nil
}
