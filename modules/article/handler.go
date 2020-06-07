package article

import (
	"base-site-api/errors"
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
		h.JSON(c, 500, &responses.ErrorResponse{
			Message: "Problem with getting the articles",
			Error:   err.Error(),
		})
		return
	}

	p := h.CalculatePagination(page, size, count)

	a := PaginatedArticles{
		p,
		articles,
	}

	h.JSON(c, 400, &a)
}

func (h *Handler) Create(c *fiber.Ctx) {

	article := &models.Article{}

	err := c.BodyParser(article)

	if err != nil {
		log.Errorf("Error while parsing article %s", err)

		h.JSON(c, 400, &responses.ErrorResponse{
			Message: "Problem with parsing the article",
			Error:   errors.ErrBadRequest.Error(),
		})
		return
	}

	a, err := h.service.Store(article, h.ParseUserId(c))

	if err != nil {
		h.JSON(c, 500, &responses.ErrorResponse{
			Message: "Problem with getting the articles",
			Error:   err.Error(),
		})
		return
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      a.ID,
	}

	h.JSON(c, 400, &r)
}

func (h *Handler) Update(c *fiber.Ctx) {
	id, err := h.ParseID(c)

	if err != nil {
		h.JSON(c, 400, &responses.ErrorResponse{
			Message: "Problem with parsing the article",
			Error:   err.Error(),
		})
		return
	}

	article := &models.Article{}

	err = c.BodyParser(article)

	if err != nil {
		h.JSON(c, 400, &responses.ErrorResponse{
			Message: "Problem with parsing the article",
			Error:   err.Error(),
		})
		return
	}

	err = h.service.Update(article, id)

	if err != nil {
		h.JSON(c, 500, &responses.ErrorResponse{
			Message: "Problem with getting the articles",
			Error:   err.Error(),
		})
		return
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      article.ID,
	}

	h.JSON(c, 400, &r)
}

func (h *Handler) Remove(c *fiber.Ctx) {
	id, err := h.ParseID(c)

	if err != nil {
		h.JSON(c, 400, &responses.ErrorResponse{
			Message: "Problem with parsing the article",
			Error:   err.Error(),
		})
		return
	}

	err = h.service.Delete(id, h.ParseUserId(c))

	if err != nil {
		h.JSON(c, 500, &responses.ErrorResponse{
			Message: "Problem with getting the articles",
			Error:   err.Error(),
		})
		return
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      uint(uID),
	}

	h.JSON(c, 400, &r)
}

func (h *Handler) GetDetail(c *fiber.Ctx) {
	slug := c.Params("slug")

	a, err := h.service.Find(slug)

	if err != nil {
		h.JSON(c, 500, &responses.ErrorResponse{
			Message: "Problem with getting the articles",
			Error:   err.Error(),
		})
		return
	}

	h.JSON(c, 400, &a)
}
