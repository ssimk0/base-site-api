package article

import (
	"base-site-api/log"
	"base-site-api/models"
	"base-site-api/modules"
	"base-site-api/responses"
	"base-site-api/utils"

	"github.com/gofiber/fiber/v2"
)

// Handler article
type Handler struct {
	modules.Handler
	service Service
}

// NewHandler return instance of Handler
func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

// List provider list of paginated active articles
func (h *Handler) List(c *fiber.Ctx) error {
	page, size := utils.ParsePagination(c)

	articles, count, err := h.service.FindAll(c.Query("sort"), page, size)

	if err != nil {
		log.Errorf("Error while getting list of articles: %s", err)
		return h.Error(c, 500)
	}

	p := h.CalculatePagination(page, size, count)

	a := PaginatedArticles{
		p,
		articles,
	}

	return h.JSON(c, 200, &a)
}

// Create handle creating article and validation
func (h *Handler) Create(c *fiber.Ctx) error {

	article := &models.Article{}

	err := c.BodyParser(article)

	if err != nil {
		log.Debugf("Error while parsing article %s", err)
		return h.Error(c, 400)
	}

	a, err := h.service.Store(article, h.ParseUserID(c))

	if err != nil {
		log.Errorf("Error while creating article: %s", err)
		return h.Error(c, 500)
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      a.ID,
	}

	return h.JSON(c, 201, &r)
}

// Update handle update article and validation
func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Error while parsing update article ID: %s", c.Params("id"))
		return h.Error(c, 400)
	}

	article := &models.Article{}

	err = c.BodyParser(article)

	if err != nil {
		log.Debugf("Error while parsing update request article: %s ", err)
		return h.Error(c, 400)
	}

	err = h.service.Update(article, id)

	if err != nil {
		log.Errorf("Error while updating article: %s", err)
		return h.Error(c, 500)
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      id,
	}

	return h.JSON(c, 200, &r)
}

// Remove handle deleting articles
func (h *Handler) Remove(c *fiber.Ctx) error {
	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Error while parsing article id remove: %s", err)
		return h.Error(c, 400)
	}

	err = h.service.Delete(id, h.ParseUserID(c))

	if err != nil {
		log.Errorf("Error while removing article: %s", err)
		return h.Error(c, 500)
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      id,
	}

	return h.JSON(c, 200, &r)
}

// GetDetail return specific article based on slug
func (h *Handler) GetDetail(c *fiber.Ctx) error {
	slug := c.Params("slug")

	a, err := h.service.Find(slug)

	if err != nil {
		log.Debugf("Error while  get detail: %s", err)
		return h.Error(c, 404)
	}

	return h.JSON(c, 200, &a)
}
