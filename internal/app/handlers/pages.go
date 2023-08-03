package handlers

import (
	"base-site-api/internal/app/dto"
	"base-site-api/internal/app/models"
	"base-site-api/internal/log"
	"base-site-api/internal/modules/page"
	"github.com/gofiber/fiber/v2"
)

// PageHandler page
type PageHandler struct {
	Handler
	repository page.Repository
}

// NewPageHandler set the repository and return instance of PageHandler
func NewPageHandler(s page.Repository) *PageHandler {
	return &PageHandler{
		repository: s,
	}
}

// ListCategories returns all page categories
func (h *PageHandler) ListCategories(c *fiber.Ctx) error {
	categories, err := h.repository.FindCategories()

	if err != nil {
		log.Errorf("Error while getting list categories %s", err)

		return h.Error(500)
	}

	return h.JSON(c, 200, categories)
}

// ListPages return all pages specific for page category
func (h *PageHandler) ListPages(c *fiber.Ctx) error {
	pages, err := h.repository.FindAllByCategorySlug(c.Params("pageCategory"))

	if err != nil {
		log.Debugf("Error while getting pages by category slug %s", err)
		return h.Error(404)
	}

	return h.JSON(c, 200, pages)
}

// GetDetail return detail for page by slug
func (h *PageHandler) GetDetail(c *fiber.Ctx) error {
	p, children, err := h.repository.FindBySlug(c.Params("slug"))

	if err != nil {
		log.Debugf("Error while getting page %s", err)
		return h.Error(404)
	}

	return h.JSON(c, 200, &page.PageDetail{
		Page:     *p,
		Children: children,
	})
}

// Create the page
func (h *PageHandler) Create(c *fiber.Ctx) error {
	p := &models.Page{}
	categorySlug := c.Params("pageCategory")

	err := c.BodyParser(p)

	if err != nil {
		log.Debugf("Error while parsing page create %s", err)
		return h.Error(400)
	}

	pageID, err := h.repository.Store(p, categorySlug, h.ParseUserID(c))

	if err != nil {
		log.Errorf("Error while create page %s", err)
		return h.Error(500)
	}

	return h.JSON(c, 201, &dto.SuccessResponse{
		Success: true,
		ID:      pageID,
	})
}

// Update page
func (h *PageHandler) Update(c *fiber.Ctx) error {
	p := &models.Page{}

	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Problem with parsing update page id: %v", c.Params("id"))
		return h.Error(400)
	}

	err = c.BodyParser(p)

	if err != nil {
		log.Debugf("Error while parsing page update %s", err)
		return h.Error(400)
	}

	err = h.repository.Update(p, id)

	if err != nil {
		log.Errorf("Error while update page %s", err)
		return h.Error(500)
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

		return h.Error(400)
	}

	err = h.repository.Delete(id, h.ParseUserID(c))

	if err != nil {
		log.Errorf("Problem while removing page: %s", err)

		return h.Error(500)
	}

	r := dto.SuccessResponse{
		Success: true,
		ID:      id,
	}

	return h.JSON(c, 200, &r)
}
