package routes

import (
	"base-site-api/internal/app/handlers"
	"base-site-api/internal/database"
	"base-site-api/internal/middleware"
	"base-site-api/internal/modules/announcement"
	"base-site-api/internal/modules/article"
	"base-site-api/internal/modules/auth"
	"base-site-api/internal/modules/page"
	"base-site-api/internal/modules/upload"
	"github.com/gofiber/fiber/v2"
)

func postOnlyFilter(c *fiber.Ctx) bool {
	return c.Method() == "POST"
}

func Register(api fiber.Router, signingKey []byte, templatePath string) {
	ah := handlers.NewArticleHandler(article.NewService(article.NewRepository(database.Instance())))
	authHandler := handlers.NewAuthHandler(auth.NewService(auth.NewRepository(database.Instance()), signingKey, templatePath))

	articles := api.Group("/v1/articles")
	articles.Use(middleware.NewAuthMiddleware(&middleware.Config{
		SigningKey: signingKey,
		Filter:     middleware.FilterGetOnly,
		DB:         database.Instance(),
	}))

	articles.Get("/", ah.List)
	articles.Post("/", ah.Create)
	articles.Put("/:id", ah.Update)
	articles.Delete("/:id", ah.Remove)
	articles.Get("/:slug", ah.GetDetail)

	a := api.Group("/v1/auth")
	a.Use(middleware.NewAuthMiddleware(&middleware.Config{
		SigningKey: signingKey,
		Filter:     postOnlyFilter,
		DB:         database.Instance(),
	}))

	a.Get("/user", authHandler.GetUserInfo)
	a.Post("/login", authHandler.Login)
	a.Post("/register-user", authHandler.RegisterUser)
	a.Post("/forgot-password", authHandler.ForgotPassword)
	a.Post("/reset-password/:token", authHandler.ResetPassword)

	ph := handlers.NewPageHandler(page.NewRepository(database.Instance()))

	pages := api.Group("/v1/pages")
	pages.Use(middleware.NewAuthMiddleware(&middleware.Config{
		SigningKey: signingKey,
		Filter:     middleware.FilterGetOnly,
		DB:         database.Instance(),
	}))

	pages.Get("/", ph.ListCategories)
	pages.Get("/:pageCategory", ph.ListPages)
	pages.Get("/:pageCategory/:slug", ph.GetDetail)
	pages.Post("/:pageCategory", ph.Create)
	pages.Put("/:id", ph.Update)
	pages.Delete("/:id", ph.Remove)

	anh := handlers.NewAnnouncementHandler(announcement.NewService(announcement.NewRepository(database.Instance())))

	announce := api.Group("/v1/announcement")
	announce.Use(middleware.NewAuthMiddleware(&middleware.Config{
		SigningKey: signingKey,
		Filter:     middleware.FilterGetOnly,
		DB:         database.Instance(),
	}))

	announce.Get("/", anh.Active)
	announce.Post("/", anh.Create)

	uh := handlers.NewUploadHandler(upload.NewService(upload.NewRepository(database.Instance())))

	uploads := api.Group("/v1/uploads")
	uploads.Use(middleware.NewAuthMiddleware(&middleware.Config{
		SigningKey: signingKey,
		Filter:     middleware.FilterGetOnly,
		DB:         database.Instance(),
	}))

	uploads.Get("/:type", uh.ListCategories)
	uploads.Get("/:type/:uploadCategory/latest", uh.LastestUpload)
	uploads.Get("/:type/:uploadCategory/:id/download", uh.DownloadUpload)
	uploads.Get("/:type/:uploadCategory/:id", uh.Detail)
	uploads.Put("/:type/:uploadCategory/:id", uh.EditUpload)
	uploads.Get("/:type/:uploadCategory", uh.ListUploads)
	uploads.Post("/:type/:uploadCategory", uh.Upload)
	uploads.Post("/:type", uh.CreateCategory)
	uploads.Put("/:type/:id", uh.UpdateCategory)
	uploads.Delete("/:type/:uploadCategory/:id", uh.Remove)
	uploads.Delete("/:type/:id", uh.RemoveCategory)
}
