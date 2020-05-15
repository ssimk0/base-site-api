package article

import (
	"base-site-api/models"
	"base-site-api/responses"

	"github.com/gofiber/fiber"
)

type ArticleHandler struct {
	service Service
}

func NewHandler(s Service) *ArticleHandler {
	return &ArticleHandler{
		service: s,
	}
}

func (h *ArticleHandler) List(c *fiber.Ctx) {
	articles, err := h.service.FindAll(c.Params("sort"))

	if err != nil {
		c.Status(500).Send(responses.ErrorResponse{
			Message: "Problem with getting the articles",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(&articles)
}

func (h *ArticleHandler) Create(c *fiber.Ctx) {
	article := &models.Article{}

	err := c.BodyParser(article)

	if err != nil {
		c.Status(400).Send(responses.ErrorResponse{
			Message: "Problem with parsing the article",
			Error:   err.Error(),
		})
		return
	}

	a, err := h.service.Store(article)

	if err != nil {
		c.Status(500).Send(responses.ErrorResponse{
			Message: "Problem with getting the articles",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(responses.SuccessResponse{
		Success: true,
		Id: a.ID,
	})
}