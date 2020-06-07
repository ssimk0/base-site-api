package page

import (
	"base-site-api/errors"
	"base-site-api/models"
	"base-site-api/modules"
	"base-site-api/responses"
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

func (h *Handler) ListCategories(c *fiber.Ctx) {
	categories, err := h.service.FindCategories()

	if err != nil {
		log.Debugf("Error while getting list categories %s", err)
		h.JSON(c, 500, responses.ErrorResponse{
			Error:   errors.InternalServerError.Error(),
			Message: "Problem while listing categories",
		})
		return
	}

	h.JSON(c, 200, categories)
}

func (h *Handler) ListPages(c *fiber.Ctx) {
	pages, err := h.service.FindAllByCategory(c.Params("page-category"))

	if err != nil {
		log.Debugf("Error while getting pages by category slug %s", err)
		h.JSON(c, 404, responses.ErrorResponse{
			Error:   errors.NotFound.Error(),
			Message: "Pages for category not found",
		})
		return
	}

	h.JSON(c, 200, pages)
}

func (h *Handler) GetDetail(c *fiber.Ctx) {
	page, err := h.service.FindBySlug(c.Params("slug"))

	if err != nil {
		log.Debugf("Error while getting page %s", err)
		h.JSON(c, 404, responses.ErrorResponse{
			Error:   errors.NotFound.Error(),
			Message: "Page not found",
		})
		return
	}

	h.JSON(c, 200, page)
}

func (h *Handler) Create(c *fiber.Ctx) {
	page := &models.Page{}
	userID := c.Locals("userID").(uint)
	categorySlug := c.Params("page-category")

	err := c.BodyParser(page)

	if err != nil {
		log.Debugf("Error while parsing page create %s", err)
		h.JSON(c, 400, responses.ErrorResponse{
			Error:   errors.BadRequest.Error(),
			Message: "Cannot parse page",
		})
		return
	}

	pageId, err := h.service.Create(page, categorySlug, userID)

	if err != nil {
		log.Debugf("Error while create page %s", err)
		h.JSON(c, 500, responses.ErrorResponse{
			Error:   errors.InternalServerError.Error(),
			Message: "Cannot create page",
		})
		return
	}

	h.JSON(c, 200, &responses.SuccessResponse{
		Success: true,
		ID:      pageId,
	})
}

func (h *Handler) Update(c *fiber.Ctx) {
	page := &models.Page{}

	err := c.BodyParser(page)

	if err != nil {
		log.Debugf("Error while parsing page update %s", err)
		h.JSON(c, 400, responses.ErrorResponse{
			Error:   errors.BadRequest.Error(),
			Message: "Cannot parse page",
		})
		return
	}

	err = h.service.Update(page, page.ID)

	if err != nil {
		log.Debugf("Error while update page %s", err)
		h.JSON(c, 500, responses.ErrorResponse{
			Error:   errors.NotFound.Error(),
			Message: "Cannot update page",
		})
		return
	}

	h.JSON(c, 200, &responses.SuccessResponse{
		Success: true,
		ID:      page.ID,
	})
}

func (h *Handler) Remove(c *fiber.Ctx) {
	userID := c.Locals("userID").(uint)

	id := c.Params("id")
	uID, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		h.JSON(c, 400, &responses.ErrorResponse{
			Message: "Problem with parsing the article",
			Error:   err.Error(),
		})
		return
	}

	err = h.service.Delete(uint(uID), userID)

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
