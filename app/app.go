package app

import (
	"base-site-api/config"
	"base-site-api/responses"

	"github.com/gofiber/fiber"
)

// NewApp function prepare whole app setup
func NewApp(c *config.Config) *fiber.App {
	// SETUP APP
	app := fiber.New(&fiber.Settings{
		Prefork:               true,
		CaseSensitive:         true,
		StrictRouting:         true,
		DisableStartupMessage: true,
		ErrorHandler: func(ctx *fiber.Ctx, err error) {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Return HTTP response
			ctx.Status(code).JSON(responses.ErrorResponse{
				Error:   err.Error(),
				Message: "",
			})
		},
	})

	configureGlobalMiddleware(app)

	configureAPIRoutes(app, c)

	return app
}
