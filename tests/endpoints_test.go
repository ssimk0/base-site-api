package tests

import (
	"base-site-api/app"
	"base-site-api/config"
	"base-site-api/modules"

	"github.com/gofiber/fiber/v2"
)

type EndpointsTestSuite struct {
	app *fiber.App
	modules.RepositoryTestSuite
}

func (s *EndpointsTestSuite) SetupApp() {
	s.Setup()

	c := config.Config{
		Constants: config.Constants{
			ADDRESS:      "127.0.0.1:8081",
			ENV:          "test",
			TemplatePath: "../templates",
		},
		Database: s.Conn,
	}

	s.app = app.NewApp(&c)
}
