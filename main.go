package main

import (
	"os"

	"base-site-api/internal/config"

	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
)

func init() {
	var logLevel log.Level

	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)

	if os.Getenv("GO_ENV") == "development" {
		logLevel = log.DebugLevel
	} else {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
}

func main() {
	config, err := config.New()

	if err != nil {
		log.Fatal(err)
	}

	// SETUP APP
	app := fiber.New(&fiber.Settings{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
	})

	configureGlobalMiddleware(app)

	configureAPIRoutes(app, config)

	startServer(app, config)
}

func startServer(app *fiber.Fiber, config *config.Config) {
	err := app.Listen(config.Constants.ADDRESS)

	if err != nil {
		log.Fatal(err.Error())
	}
}
