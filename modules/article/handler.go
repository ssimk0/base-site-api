package article

import (
	"base-site-api/models"
	"base-site-api/modules"
	"base-site-api/responses"
	"base-site-api/utils"
	"github.com/gofiber/fiber"
	"strconv"
)

type ArticleHandler struct {
	modules.Handler
	service Service
}

func NewHandler(s Service) *ArticleHandler {
	return &ArticleHandler{
		service: s,
	}
}

func (h *ArticleHandler) List(c *fiber.Ctx) {
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

	c.JSON(&a)
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
	slug := c.Params("slug")

	article, err := h.service.Find(slug)

	if err != nil {
		c.Status(500).Send(responses.ErrorResponse{
			Message: "Problem with getting the articles",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(article)
}
