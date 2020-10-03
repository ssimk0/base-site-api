package http

import (
	"base-site-api/internal/dto"

	"github.com/gofiber/fiber/v2"
)

// NewApp function prepare whole app setup
func NewApp(c *Config) *fiber.App {
	// SETUP APP
	app := fiber.New(fiber.Config{
		Prefork:               true,
		CaseSensitive:         true,
		StrictRouting:         true,
		DisableStartupMessage: true,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Return HTTP response
			return c.Status(code).JSON(dto.ErrorResponse{
				Error:   err.Error(),
				Message: "",
			})
		},
	})

	configureGlobalMiddleware(app)

	configureAPIRoutes(app, c)

	return app
}
