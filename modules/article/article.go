package article

import (
	"base-site-api/config"
	"base-site-api/middleware/auth"
	"github.com/gofiber/fiber"
)

type Article struct {
}

// New configure module and register all routes
func (m Article) New(config *config.Config, api *fiber.Router) {
	handler := NewHandler(NewService(NewRepository(config.Database)))

	articles := (*api).Group("/v1/articles")
	articles.Use(auth.New(&auth.Config{
		SigningKey: config.SigningKey,
		Filter:     auth.FilterGetOnly,
		DB:         config.Database,
	}))

	articles.Get("/", handler.List)
	articles.Post("/", handler.Create)
	articles.Put("/:id", handler.Update)
	articles.Delete("/:id", handler.Remove)
	articles.Get("/:slug", handler.GetDetail)
}
