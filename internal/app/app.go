package app

import (
	"base-site-api/internal/app/config"
	"base-site-api/internal/app/dto"
	"base-site-api/internal/app/models"
	"base-site-api/internal/database"
	"base-site-api/internal/email"
	"base-site-api/internal/log"
	"base-site-api/internal/routes"
	"base-site-api/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)

// New function prepare whole app setup
func New(c *config.Config) *fiber.App {
	c.Fiber.ErrorHandler = func(c *fiber.Ctx, err error) error {
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
	}

	database.Connect(&c.Database)
	// TODO: find better way to share storage and email instance
	storage.Initialize(&c.Storage)
	email.Initialize(&c.Email)

	log.Setup(&c.App)

	database.Instance().AutoMigrate(
		&models.PageCategory{},
		&models.Page{},
		&models.UploadType{},
		&models.UploadCategory{},
		&models.Upload{},
		&models.User{},
		&models.ForgotPasswordToken{},
		&models.Article{},
	)

	// SETUP APP
	app := fiber.New(c.Fiber)

	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Locals("APP_URL", c.App.AppURL)
		return ctx.Next()
	})

	// Use the Logger Middleware if enabled
	if c.Enabled["logger"] {
		app.Use(logger.New(c.Logger))
	}

	// Use the Recover Middleware if enabled
	if c.Enabled["recover"] {
		app.Use(recover.New(c.Recover))
	}
	if c.Enabled["compression"] {
		app.Use(compress.New(c.Compression))
	}

	// Use the CORS Middleware if enabled
	if c.Enabled["cors"] {
		app.Use(cors.New(c.CORS))
	}

	// Use the Helmet Middleware if enabled
	if c.Enabled["helmet"] {
		app.Use(helmet.New(c.Helmet))
	}

	api := app.Group("/api")

	routes.Register(&api, c.App.SigningKey, c.App.TemplatePath)

	return app
}
