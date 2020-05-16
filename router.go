package main

import (
	"os"

	"base-site-api/config"
	"base-site-api/modules/article"
	"base-site-api/modules/auth"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

// ENDPOINTS
func setupV1ApiEndpoints(api *fiber.Group, config *config.Config) {
	article.New(config, api)
	auth.New(config, api)
}

// SETTINGS FOR GROUPS
func configureAPIRoutes(app *fiber.Fiber, config *config.Config) {
	api := app.Group("/api", cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("ALLOWED_ORIGIN")},
	}))

	setupV1ApiEndpoints(api, config)
}
