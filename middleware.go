package main

import (
	//"base-site-api/middleware"

	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
	log "github.com/sirupsen/logrus"
)

func configureGlobalMiddleware(app *fiber.App) {
	app.Use(recover.New(recover.Config{
		Handler: func(c *fiber.Ctx, err error) {
			log.Error(err)
			c.SendStatus(500)
		},
		Log: true,
	}))
	app.Use(logger.New())
	app.Use(helmet.New())
	//app.Use(middleware.New())
}
