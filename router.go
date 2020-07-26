package main

import (
	"base-site-api/modules/page"
	"base-site-api/modules/uploads"
	"os"

	"base-site-api/config"
	"base-site-api/modules/article"
	"base-site-api/modules/auth"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

// ENDPOINTS
func setupV1ApiEndpoints(api *fiber.Router, config *config.Config) {
	article.New(config, api)
	auth.New(config, api)
	page.New(config, api)
	uploads.New(config, api)
}

// SETTINGS FOR GROUPS
func configureAPIRoutes(app *fiber.App, config *config.Config) {
	api := app.Group("/api", cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("ALLOWED_ORIGIN")},
	}))

	setupV1ApiEndpoints(&api, config)
}
