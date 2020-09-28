package uploads

import (
	"base-site-api/config"
	"base-site-api/middleware/auth"
	"github.com/gofiber/fiber/v2"
)

type Uploads struct {
}

// New prepare whole module and connect it with App
func (m Uploads) New(config *config.Config, api *fiber.Router) {
	handler := NewHandler(NewService(NewRepository(config.Database)))

	pages := (*api).Group("/v1/uploads")
	pages.Use(auth.New(&auth.Config{
		SigningKey: config.SigningKey,
		Filter:     auth.FilterGetOnly,
		DB:         config.Database,
	}))

	pages.Get("/:type", handler.ListCategories)
	pages.Get("/:type/:uploadCategory", handler.ListUploads)
	pages.Post("/:type/:uploadCategory", handler.Upload)
	pages.Post("/:type", handler.CreateCategory)
	pages.Put("/:type/:id", handler.UpdateCategory)
	pages.Delete("/:type/:uploadCategory/:id", handler.Remove)
	pages.Delete("/:type/:id", handler.RemoveCategory)
}
