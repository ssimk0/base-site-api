package uploads

import (
	"base-site-api/modules"
	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
)

// Handler for the uploads
type Handler struct {
	modules.Handler
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) ListCategories(c *fiber.Ctx) {
	t := c.Params("type")

	categories, err := h.service.UploadCategories(t)

	if err != nil {
		log.Debugf("Error while getting pages by category slug %s", err)
		h.Error(c, 404)

		return
	}

	h.JSON(c, 200, categories)
}
