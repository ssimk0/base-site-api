package main

import (
	"os"

	"base-site-api/internal/config"
	"base-site-api/modules/article"
	"base-site-api/modules/auth"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

// ENDPOINTS
func setupV1ApiEndpoints(api *fiber.Group, config *config.Config) {
	db := config.Database

	article.New(db, api)
	auth.New(db, api)
}

// SETTINGS FOR GROUPS
func configureAPIRoutes(app *fiber.Fiber, config *config.Config) {
	cors := cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("ALLOWED_ORIGIN")},
	})

	api := app.Group("/api", cors)

	setupV1ApiEndpoints(api, config)
}
