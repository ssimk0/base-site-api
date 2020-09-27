package app

import (
	"base-site-api/modules/auth"
	"base-site-api/modules/page"
	"base-site-api/modules/uploads"
	"os"

	"base-site-api/config"
	"base-site-api/modules/article"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Module is interface of api module for plug and play system to make more esier to integrate them
type Module interface {
	New(config *config.Config, api *fiber.Router)
}

// ENDPOINTS
func setupV1ApiEndpoints(api *fiber.Router, config *config.Config) {
	modules := []Module{
		article.Article{},
		auth.Auth{},
		page.Pages{},
		uploads.Uploads{},
	}

	for _, module := range modules {
		module.New(config, api)
	}
}

// SETTINGS FOR GROUPS
func configureAPIRoutes(app *fiber.App, config *config.Config) {
	api := app.Group("/api", cors.New(cors.Config{
		AllowOrigins: os.Getenv("ALLOWED_ORIGIN"),
	}))

	setupV1ApiEndpoints(&api, config)
}
