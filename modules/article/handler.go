package article

import (
	"base-site-api/errors"
	"base-site-api/models"
	"base-site-api/modules"
	"base-site-api/responses"
	"base-site-api/utils"
	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
	"strconv"
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

func (h *Handler) List(c *fiber.Ctx) {
	page, size := utils.ParsePagination(c)

	articles, count, err := h.service.FindAll(c.Query("sort"), page, size)

	if err != nil {
		c.Status(500).Send(responses.ErrorResponse{
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

	if err := c.JSON(&a); err != nil {
		c.Next(err)
	}
}

func (h *Handler) Create(c *fiber.Ctx) {
	userID := c.Locals("userID").(uint)

	article := &models.Article{}

	err := c.BodyParser(article)

	if err != nil {
		log.Errorf("Error while parsing article %s", err)

		c.Status(400).Send(responses.ErrorResponse{
			Message: "Problem with parsing the article",
			Error:   errors.BadRequest.Error(),
		})
		return
	}

	a, err := h.service.Store(article, userID)

	if err != nil {
		c.Status(500).Send(responses.ErrorResponse{
			Message: "Problem with getting the articles",
			Error:   err.Error(),
		})
		return
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      a.ID,
	}

	if err := c.JSON(&r); err != nil {
		c.Next(err)
	}
}

func (h *Handler) Update(c *fiber.Ctx) {
	article := &models.Article{}

	err := c.BodyParser(article)

	if err != nil {
		c.Status(400).Send(responses.ErrorResponse{
			Message: "Problem with parsing the article",
			Error:   err.Error(),
		})
		return
	}

	err = h.service.Update(article, article.ID)

	if err != nil {
		c.Status(500).Send(responses.ErrorResponse{
			Message: "Problem with getting the articles",
			Error:   err.Error(),
		})
		return
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      article.ID,
	}

	if err := c.JSON(&r); err != nil {
		c.Next(err)
	}
}

func (h *Handler) Remove(c *fiber.Ctx) {
	userID := c.Locals("userID").(uint)

	id := c.Params("id")
	uID, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		c.Status(400).Send(responses.ErrorResponse{
			Message: "Problem with parsing the article",
			Error:   err.Error(),
		})
		return
	}

	err = h.service.Delete(uint(uID), userID)

	if err != nil {
		c.Status(500).Send(responses.ErrorResponse{
			Message: "Problem with getting the articles",
			Error:   err.Error(),
		})
		return
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      uint(uID),
	}

	if err := c.JSON(&r); err != nil {
		c.Next(err)
	}
}

func (h *Handler) GetDetail(c *fiber.Ctx) {
	slug := c.Params("slug")

	a, err := h.service.Find(slug)

	if err != nil {
		c.Status(500).Send(responses.ErrorResponse{
			Message: "Problem with getting the articles",
			Error:   err.Error(),
		})
		return
	}

	if err := c.JSON(&a); err != nil {
		c.Next(err)
	}
}
