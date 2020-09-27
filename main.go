package main

import (
	"base-site-api/app"
	"base-site-api/config"
	"base-site-api/log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// log.SetupLogger()

	c, err := config.New()

	if err != nil {
		log.Fatal(err)
	}

	app := app.NewApp(c)
	startServer(app, c)
}

func startServer(app *fiber.App, c *config.Config) {

	err := app.Listen(c.Constants.ADDRESS)

	if err != nil {
		log.Fatal(err)
	}
}
