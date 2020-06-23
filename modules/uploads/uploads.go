package uploads

import (
	"base-site-api/config"
	"base-site-api/middleware/auth"
	"github.com/gofiber/fiber"
)

// New prepare whole module and connect it with App
func New(config *config.Config, api *fiber.Group) {
	handler := NewHandler(NewService(NewRepository(config.Database)))

	pages := api.Group("/v1/uploads")
	pages.Use(auth.New(&auth.Config{
		SigningKey: config.SigningKey,
		Filter:     auth.FilterGetOnly,
	}))

	pages.Get("/:type", handler.ListCategories)
	pages.Get("/:type/:upload-category", handler.ListUploads)
	pages.Post("/:type/:upload-category", handler.Upload)
	//pages.Post("/:page-category", handler.Create)
	//pages.Put("/:id", handler.Update)
	//pages.Delete("/:id", handler.Remove)

}
