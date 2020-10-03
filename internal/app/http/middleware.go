package http

import (
	//"base-site-api/middleware"

	"base-site-api/internal/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

// TODO: wrote middleware for user permission
func configureGlobalMiddleware(app *fiber.App) {
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Output: log.Writer(),
	}))
	app.Use(helmet.New())
}
