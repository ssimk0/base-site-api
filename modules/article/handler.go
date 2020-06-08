package article

import (
	"base-site-api/models"
	"base-site-api/modules"
	"base-site-api/responses"
	"base-site-api/utils"
	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
)

// Handler article
type Handler struct {
	modules.Handler
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) List(c *fiber.Ctx) {
	page, size := utils.ParsePagination(c)

	articles, count, err := h.service.FindAll(c.Query("sort"), page, size)

	if err != nil {
		log.Errorf("Error while getting list of articles: %s", err)
		h.Error(c, 500)

		return
	}

	p := h.CalculatePagination(page, size, count)

	a := PaginatedArticles{
		p,
		articles,
	}

	h.JSON(c, 200, &a)
}

func (h *Handler) Create(c *fiber.Ctx) {

	article := &models.Article{}

	err := c.BodyParser(article)

	if err != nil {
		log.Debugf("Error while parsing article %s", err)
		h.Error(c, 400)

		return
	}

	a, err := h.service.Store(article, h.ParseUserId(c))

	if err != nil {
		log.Errorf("Error while creating article: %s", err)
		h.Error(c, 500)

		return
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      a.ID,
	}

	h.JSON(c, 201, &r)
}

func (h *Handler) Update(c *fiber.Ctx) {
	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Error while parsing update article ID: %s", c.Params("id"))
		h.Error(c, 400)

		return
	}

	article := &models.Article{}

	err = c.BodyParser(article)

	if err != nil {
		log.Debugf("Error while parsing update request article: %s ", err)
		h.Error(c, 400)

		return
	}

	err = h.service.Update(article, id)

	if err != nil {
		log.Errorf("Error while updating article: %s", err)
		h.Error(c, 500)

		return
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      id,
	}

	h.JSON(c, 200, &r)
}

func (h *Handler) Remove(c *fiber.Ctx) {
	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Error while parsing article id remove: %s", err)
		h.Error(c, 400)

		return
	}

	err = h.service.Delete(id, h.ParseUserId(c))

	if err != nil {
		log.Errorf("Error while removing article: %s", err)
		h.Error(c, 500)

		return
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      id,
	}

	h.JSON(c, 200, &r)
}

func (h *Handler) GetDetail(c *fiber.Ctx) {
	slug := c.Params("slug")

	a, err := h.service.Find(slug)

	if err != nil {
		log.Debugf("Error while  get detail: %s", err)
		h.Error(c, 404)

		return
	}

	h.JSON(c, 200, &a)
}
