package integration_tests

import (
	"base-site-api/internal/app"
	"base-site-api/internal/app/config"
	"base-site-api/internal/database"
	"base-site-api/internal/tests/test_helper"
	"github.com/gofiber/fiber/v2"
)

type EndpointsTestSuite struct {
	app *fiber.App
	test_helper.RepositoryTestSuite
}

func (s *EndpointsTestSuite) SetupApp() {
	s.Setup()

	c := config.Config{
		App: config.ApplicationConfiguration{
			Listen:       "127.0.0.1:8081",
			Debug:        true,
			LogToFile:    false,
			TemplatePath: "../../../templates",
		},
		Database: config.DatabaseConfiguration{
			Database: "/tmp/test.db",
			Driver:   "sqlite",
		},
	}

	s.app = app.New(&c)
	s.Conn = database.Instance()
}
