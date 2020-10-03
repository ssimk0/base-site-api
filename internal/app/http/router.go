package http

import (
	"base-site-api/internal/article"
	"base-site-api/internal/auth"
	"base-site-api/internal/middleware"
	"base-site-api/internal/page"
	"base-site-api/internal/transfer/http"
	"base-site-api/internal/upload"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func postOnlyFilter(c *fiber.Ctx) bool {
	return c.Method() == "POST"
}

// ENDPOINTS
func setupV1ApiEndpoints(api *fiber.Router, config *Config) {
	ah := http.NewArticleHandler(article.NewService(article.NewRepository(config.Database)))
	authHandler := http.NewAuthHandler(auth.NewService(auth.NewRepository(config.Database), config.SigningKey, config.Constants.TemplatePath))

	articles := (*api).Group("/v1/articles")
	articles.Use(middleware.NewAuthMiddleware(&middleware.Config{
		SigningKey: config.SigningKey,
		Filter:     middleware.FilterGetOnly,
		DB:         config.Database,
	}))

	articles.Get("/", ah.List)
	articles.Post("/", ah.Create)
	articles.Put("/:id", ah.Update)
	articles.Delete("/:id", ah.Remove)
	articles.Get("/:slug", ah.GetDetail)

	a := (*api).Group("/v1/auth")
	a.Use(middleware.NewAuthMiddleware(&middleware.Config{
		SigningKey: config.SigningKey,
		Filter:     postOnlyFilter,
		DB:         config.Database,
	}))

	a.Get("/user", authHandler.GetUserInfo)
	a.Post("/login", authHandler.Login)
	a.Post("/register-user", authHandler.RegisterUser)
	a.Post("/forgot-password", authHandler.ForgotPassword)
	a.Post("/reset-password/:token", authHandler.ResetPassword)

	ph := http.NewPageHandler(page.NewService(page.NewRepository(config.Database)))

	pages := (*api).Group("/v1/pages")
	pages.Use(middleware.NewAuthMiddleware(&middleware.Config{
		SigningKey: config.SigningKey,
		Filter:     middleware.FilterGetOnly,
		DB:         config.Database,
	}))

	pages.Get("/", ph.ListCategories)
	pages.Get("/:pageCategory", ph.ListPages)
	pages.Get("/:pageCategory/:slug", ph.GetDetail)
	pages.Post("/:pageCategory", ph.Create)
	pages.Put("/:id", ph.Update)
	pages.Delete("/:id", ph.Remove)

	uh := http.NewUploadHandler(upload.NewService(upload.NewRepository(config.Database)))

	uploads := (*api).Group("/v1/uploads")
	uploads.Use(middleware.NewAuthMiddleware(&middleware.Config{
		SigningKey: config.SigningKey,
		Filter:     middleware.FilterGetOnly,
		DB:         config.Database,
	}))

	uploads.Get("/:type", uh.ListCategories)
	uploads.Get("/:type/:uploadCategory", uh.ListUploads)
	uploads.Post("/:type/:uploadCategory", uh.Upload)
	uploads.Post("/:type", uh.CreateCategory)
	uploads.Put("/:type/:id", uh.UpdateCategory)
	uploads.Delete("/:type/:uploadCategory/:id", uh.Remove)
	uploads.Delete("/:type/:id", uh.RemoveCategory)
}

// SETTINGS FOR GROUPS
func configureAPIRoutes(app *fiber.App, config *Config) {
	api := app.Group("/api", cors.New(cors.Config{
		AllowOrigins: os.Getenv("ALLOWED_ORIGIN"),
	}))

	setupV1ApiEndpoints(&api, config)
}
