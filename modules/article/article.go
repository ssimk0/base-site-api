package article

import (
	"base-site-api/modules/article/delivery/http"
	"base-site-api/modules/article/repository"
	"base-site-api/modules/article/service"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func New(db *gorm.DB, api *fiber.Group) {
	handler := http.New(service.New(repository.New(db)))

	api.Get("/articles", handler.List)
	// api.Post("/articles", handler.Create)
	// api.Put("/articles", handler.Update)
	// api.Delete("/articles/{id}", handler.Remove)
	// api.Get("/articles/{id}", handler.GetDetail)
}
