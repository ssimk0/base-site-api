package integration_tests

import (
	"base-site-api/internal/app/http"
	"base-site-api/internal/common"
	"github.com/gofiber/fiber/v2"
)

type EndpointsTestSuite struct {
	app *fiber.App
	common.RepositoryTestSuite
}

func (s *EndpointsTestSuite) SetupApp() {
	s.Setup()

	c := http.Config{
		Constants: http.Constants{
			ADDRESS:      "127.0.0.1:8081",
			ENV:          "test",
			TemplatePath: "../../templates",
		},
		Database: s.Conn,
	}

	s.app = http.NewApp(&c)
}
