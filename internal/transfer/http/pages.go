package http

import (
	"base-site-api/internal/common"
	"base-site-api/internal/dto"
	"base-site-api/internal/log"
	"base-site-api/internal/page"
	"github.com/gofiber/fiber/v2"
)

// PageHandler page
type PageHandler struct {
	common.Handler
	service page.Service
}

// NewHandler set  the service and return instance  of PageHandler
func NewPageHandler(s page.Service) *PageHandler {
	return &PageHandler{
		service: s,
	}
}

// ListCategories returns all page categories
func (h *PageHandler) ListCategories(c *fiber.Ctx) error {
	categories, err := h.service.FindCategories()

	if err != nil {
		log.Errorf("Error while getting list categories %s", err)

		return h.Error(c, 500)
	}

	return h.JSON(c, 200, categories)
}

// ListPages return all pages specific for page category
func (h *PageHandler) ListPages(c *fiber.Ctx) error {
	pages, err := h.service.FindAllByCategory(c.Params("pageCategory"))

	if err != nil {
		log.Debugf("Error while getting pages by category slug %s", err)
		return h.Error(c, 404)

	}

	return h.JSON(c, 200, pages)
}

// GetDetail return detail for page by slug
func (h *PageHandler) GetDetail(c *fiber.Ctx) error {
	page, err := h.service.FindBySlug(c.Params("slug"))

	if err != nil {
		log.Debugf("Error while getting page %s", err)
		return h.Error(c, 404)
	}

	return h.JSON(c, 200, page)
}

// Create the page
func (h *PageHandler) Create(c *fiber.Ctx) error {
	page := &page.Page{}
	categorySlug := c.Params("pageCategory")

	err := c.BodyParser(page)

	if err != nil {
		log.Debugf("Error while parsing page create %s", err)
		return h.Error(c, 400)
	}

	pageID, err := h.service.Store(page, categorySlug, h.ParseUserID(c))

	if err != nil {
		log.Errorf("Error while create page %s", err)
		return h.Error(c, 500)
	}

	return h.JSON(c, 201, &dto.SuccessResponse{
		Success: true,
		ID:      pageID,
	})
}

// Update page
func (h *PageHandler) Update(c *fiber.Ctx) error {
	page := &page.Page{}

	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Problem with parsing update page id: %v", c.Params("id"))
		return h.Error(c, 400)
	}

	err = c.BodyParser(page)

	if err != nil {
		log.Debugf("Error while parsing page update %s", err)
		return h.Error(c, 400)
	}

	err = h.service.Update(page, id)

	if err != nil {
		log.Errorf("Error while update page %s", err)
		return h.Error(c, 500)
	}

	return h.JSON(c, 200, &dto.SuccessResponse{
		Success: true,
		ID:      id,
	})
}

// Remove page
func (h *PageHandler) Remove(c *fiber.Ctx) error {
	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Problem with parsing page id for remove: %v", c.Params("id"))

		return h.Error(c, 400)
	}

	err = h.service.Delete(id, h.ParseUserID(c))

	if err != nil {
		log.Errorf("Problem while removing page: %s", err)

		return h.Error(c, 500)
	}

	r := dto.SuccessResponse{
		Success: true,
		ID:      id,
	}

	return h.JSON(c, 200, &r)
}
