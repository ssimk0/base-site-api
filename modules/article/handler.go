package article

import (
	"base-site-api/models"
	"base-site-api/responses"
	"github.com/gofiber/fiber"
	"strconv"
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
	articles, err := h.service.FindAll(c.Query("sort"))

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
	userID := c.Locals("userID").(uint)

	article := &models.Article{}

	err := c.BodyParser(article)

	if err != nil {
		c.Status(400).Send(responses.ErrorResponse{
			Message: "Problem with parsing the article",
			Error:   err.Error(),
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

	c.JSON(responses.SuccessResponse{
		Success: true,
		Id:      a.ID,
	})
}

func (h *ArticleHandler) Update(c *fiber.Ctx) {
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

	c.JSON(responses.SuccessResponse{
		Success: true,
		Id:      article.ID,
	})
}

func (h *ArticleHandler) Remove(c *fiber.Ctx) {
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

	c.JSON(responses.SuccessResponse{
		Success: true,
		Id:      uint(uID),
	})
}

func (h *ArticleHandler) GetDetail(c *fiber.Ctx) {
	id := c.Params("id")
	uID, err := strconv.ParseUint(id, 16, 32)

	if err != nil {
		c.Status(400).Send(responses.ErrorResponse{
			Message: "Problem with parsing the article",
			Error:   err.Error(),
		})
		return
	}

	article, err := h.service.Find(uint(uID))

	if err != nil {
		c.Status(500).Send(responses.ErrorResponse{
			Message: "Problem with getting the articles",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(article)
}
