package app

import (
	//"base-site-api/middleware"

	"base-site-api/log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
)

// TODO: wrote middleware for user permission
func configureGlobalMiddleware(app *fiber.App) {
	app.Use(recover.New(recover.Config{
		Handler: func(c *fiber.Ctx, err error) {
			log.Error(err)
			c.SendStatus(500)
		},
		Log: true,
	}))
	app.Use(logger.New(logger.Config{
		Output: log.Writer(),
	}))
	app.Use(helmet.New())
}
