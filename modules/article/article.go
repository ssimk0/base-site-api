package article

import (
	"base-site-api/middleware"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func New(db *gorm.DB, api *fiber.Group) {
	handler := NewHandler(NewService(NewRepository(db)))
	articles := api.Group("/v1/articles")
	articles.Use(middleware.New())	

	articles.Get("/", handler.List)
	// api.Post("/v1/articles", handler.Create)
	// api.Put("/v1/articles", handler.Update)
	// api.Delete("/v1/articles/{id}", handler.Remove)
	// api.Get("/v1/articles/{id}", handler.GetDetail)
}
