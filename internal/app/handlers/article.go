package handlers

import (
	"base-site-api/internal/app/dto"
	"base-site-api/internal/app/models"
	"base-site-api/internal/log"
	"base-site-api/internal/modules/article"
	"base-site-api/internal/pagination"
	"github.com/gofiber/fiber/v2"
)

// ArticleHandler article
type ArticleHandler struct {
	Handler
	service article.Service
}

// NewHandler return instance of ArticleHandler
func NewArticleHandler(s article.Service) *ArticleHandler {
	return &ArticleHandler{
		service: s,
	}
}

// List provider list of paginated active articles
func (h *ArticleHandler) List(c *fiber.Ctx) error {
	page, size := pagination.ParsePagination(c)

	articles, count, err := h.service.FindAll(c.Query("sort"), page, size)

	if err != nil {
		log.Errorf("Error while getting list of articles: %s", err)
		return h.Error(500)
	}

	p := h.CalculatePagination(page, size, count)

	a := article.PaginatedArticles{
		Pagination: p,
		Articles:   articles,
	}

	return h.JSON(c, 200, &a)
}

// Create handle creating article and validation
func (h *ArticleHandler) Create(c *fiber.Ctx) error {

	data := &models.Article{}

	err := c.BodyParser(data)

	if err != nil {
		log.Debugf("Error while parsing article %s", err)
		return h.Error(400)
	}

	a, err := h.service.Store(data, h.ParseUserID(c))

	if err != nil {
		log.Errorf("Error while creating article: %s", err)
		return h.Error(500)
	}

	r := dto.SuccessResponse{
		Success: true,
		ID:      a.ID,
	}

	return h.JSON(c, 201, &r)
}

// Update handle update article and validation
func (h *ArticleHandler) Update(c *fiber.Ctx) error {
	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Error while parsing update article ID: %s", c.Params("id"))
		return h.Error(400)
	}

	data := &models.Article{}

	err = c.BodyParser(data)

	if err != nil {
		log.Debugf("Error while parsing update request article: %s ", err)
		return h.Error(400)
	}

	err = h.service.Update(data, id)

	if err != nil {
		log.Errorf("Error while updating article: %s", err)
		return h.Error(500)
	}

	r := dto.SuccessResponse{
		Success: true,
		ID:      id,
	}

	return h.JSON(c, 200, &r)
}

// Remove handle deleting articles
func (h *ArticleHandler) Remove(c *fiber.Ctx) error {
	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Error while parsing article id remove: %s", err)
		return h.Error(400)
	}

	err = h.service.Delete(id, h.ParseUserID(c))

	if err != nil {
		log.Errorf("Error while removing article: %s", err)
		return h.Error(500)
	}

	r := dto.SuccessResponse{
		Success: true,
		ID:      id,
	}

	return h.JSON(c, 200, &r)
}

// GetDetail return specific article based on slug
func (h *ArticleHandler) GetDetail(c *fiber.Ctx) error {
	slug := c.Params("slug")

	a, err := h.service.Find(slug)

	if err != nil {
		log.Debugf("Error while  get detail: %s", err)
		return h.Error(404)
	}

	return h.JSON(c, 200, &a)
}
