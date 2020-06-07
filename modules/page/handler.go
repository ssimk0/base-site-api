package page

import (
	"base-site-api/errors"
	"base-site-api/modules"
	"base-site-api/responses"
	"github.com/gofiber/fiber"
)

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
	categories, err := h.service.FindCategories()

	if err != nil {
		h.JSON(c, 500, responses.ErrorResponse{
			Error:   errors.InternalServerError.Error(),
			Message: "Problem while listing categories",
		})
	}

	h.JSON(c, 200, categories)
}

func (h *Handler) ListPages(c *fiber.Ctx) {
	pages, err := h.service.FindAllByCategory(c.Params("page-category"))

	if err != nil {
		h.JSON(c, 404, responses.ErrorResponse{
			Error:   errors.NotFound.Error(),
			Message: "Pages for category not found",
		})
	}

	h.JSON(c, 200, pages)
}
