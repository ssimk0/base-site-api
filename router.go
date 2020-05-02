package main

import (
	"os"

	"base-site-api/internal/config"
	"base-site-api/modules/article"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

// ENDPOINTS
func setupV1ApiEndpoints(api *fiber.Group, config *config.Config) {
	v1 := api.Group("/v1")
	db := config.Database

	article.New(db, v1)
}

// SETTINGS FOR GROUPS
func configureAPIRoutes(app *fiber.Fiber, config *config.Config) {
	cors := cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("ALLOWED_ORIGIN")},
	})

	api := app.Group("/api", cors)

	setupV1ApiEndpoints(api, config)
}
