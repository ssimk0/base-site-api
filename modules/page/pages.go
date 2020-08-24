package page

import (
	"base-site-api/config"
	"base-site-api/middleware/auth"

	"github.com/gofiber/fiber"
)

type Pages struct {
}

// New prepare whole module and connect it with App
func (m Pages) New(config *config.Config, api *fiber.Router) {
	handler := NewHandler(NewService(NewRepository(config.Database)))

	pages := (*api).Group("/v1/pages")
	pages.Use(auth.New(&auth.Config{
		SigningKey: config.SigningKey,
		Filter:     auth.FilterGetOnly,
		DB:         config.Database,
	}))

	pages.Get("/", handler.ListCategories)
	pages.Get("/:pageCategory", handler.ListPages)
	pages.Get("/:pageCategory/:slug", handler.GetDetail)
	pages.Post("/:pageCategory", handler.Create)
	pages.Put("/:id", handler.Update)
	pages.Delete("/:id", handler.Remove)

}
