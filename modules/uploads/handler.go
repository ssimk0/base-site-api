package uploads

import (
	"base-site-api/modules"
	"base-site-api/utils"
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
	s := c.Params("type")

	categories, err := h.service.UploadCategories(s)

	if err != nil {
		log.Debugf("Error while getting upload categories by type slug %s", err)
		h.Error(c, 404)

		return
	}

	h.JSON(c, 200, categories)
}

func (h *Handler) ListUploads(c *fiber.Ctx) {
	s := c.Params("upload-category")
	page, size := utils.ParsePagination(c)

	uploads, count, err := h.service.UploadsByCategory(s, page, size)

	if err != nil {
		log.Debugf("Error while getting upload by category slug %s", err)
		h.Error(c, 404)

		return
	}

	p := h.CalculatePagination(page, size, count)

	h.JSON(c, 200, PaginatedUploads{
		p,
		uploads,
	})
}

func (h *Handler) Upload(c *fiber.Ctx) {
	file, err := c.FormFile("file")
	s := c.Params("upload-category")
	t := c.Params("type")

	if err != nil {
		log.Debugf("Error while parsing upload %s", err)
		h.Error(c, 400)

		return
	}

	r, err := h.service.Store(file, s, t)

	if err != nil {
		log.Debugf("Error while upload %s", err)
		h.Error(c, 400)

		return
	}

	h.JSON(c, 200, r)
}
