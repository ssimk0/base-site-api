package page

import (
	"base-site-api/models"
	"base-site-api/modules"
	"base-site-api/responses"
	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
)

// Handler page
type Handler struct {
	modules.Handler
	service Service
}

// NewHandler set  the service and return instance  of Handler
func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

// ListCategories returns all page categories
func (h *Handler) ListCategories(c *fiber.Ctx) {
	categories, err := h.service.FindCategories()

	if err != nil {
		log.Errorf("Error while getting list categories %s", err)
		h.Error(c, 500)

		return
	}

	h.JSON(c, 200, categories)
}

// ListPages return all pages specific for page category
func (h *Handler) ListPages(c *fiber.Ctx) {
	pages, err := h.service.FindAllByCategory(c.Params("page-category"))

	if err != nil {
		log.Debugf("Error while getting pages by category slug %s", err)
		h.Error(c, 404)

		return
	}

	h.JSON(c, 200, pages)
}

// GetDetail return detail for page by slug
func (h *Handler) GetDetail(c *fiber.Ctx) {
	page, err := h.service.FindBySlug(c.Params("slug"))

	if err != nil {
		log.Debugf("Error while getting page %s", err)
		h.Error(c, 404)

		return
	}

	h.JSON(c, 200, page)
}

// Create the page
func (h *Handler) Create(c *fiber.Ctx) {
	page := &models.Page{}
	categorySlug := c.Params("page-category")

	err := c.BodyParser(page)

	if err != nil {
		log.Debugf("Error while parsing page create %s", err)
		h.Error(c, 400)

		return
	}

	pageID, err := h.service.Store(page, categorySlug, h.ParseUserID(c))

	if err != nil {
		log.Errorf("Error while create page %s", err)
		h.Error(c, 500)

		return
	}

	h.JSON(c, 201, &responses.SuccessResponse{
		Success: true,
		ID:      pageID,
	})
}

// Update page
func (h *Handler) Update(c *fiber.Ctx) {
	page := &models.Page{}

	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Problem with parsing update page id: %v", c.Params("id"))
		h.Error(c, 400)

		return
	}

	err = c.BodyParser(page)

	if err != nil {
		log.Debugf("Error while parsing page update %s", err)
		h.Error(c, 400)

		return
	}

	err = h.service.Update(page, id)

	if err != nil {
		log.Errorf("Error while update page %s", err)
		h.Error(c, 500)

		return
	}

	h.JSON(c, 200, &responses.SuccessResponse{
		Success: true,
		ID:      id,
	})
}

// Remove page
func (h *Handler) Remove(c *fiber.Ctx) {
	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Problem with parsing page id for remove: %v", c.Params("id"))
		h.Error(c, 400)

		return
	}

	err = h.service.Delete(id, h.ParseUserID(c))

	if err != nil {
		log.Errorf("Problem while removing page: %s", err)
		h.Error(c, 500)

		return
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      id,
	}

	h.JSON(c, 200, &r)
}
