package article

import (
	"base-site-api/internal/config"
	"base-site-api/middleware/auth"
	"github.com/gofiber/fiber"
)

func New(config *config.Config, api *fiber.Group) {
	handler := NewHandler(NewService(NewRepository(config.Database)))

	articles := api.Group("/v1/articles")
	articles.Use(auth.New(&auth.Config{
		SigningKey: config.SigningKey,
		Filter: auth.FilterOutGet,
	}))

	articles.Get("/", handler.List)
	articles.Post("/", handler.Create)
	// api.Put("/v1/articles", handler.Update)
	// api.Delete("/v1/articles/{id}", handler.Remove)
	// api.Get("/v1/articles/{id}", handler.GetDetail)
}
