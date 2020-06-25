package uploads

import (
	"base-site-api/models"
	"base-site-api/modules"
	"base-site-api/responses"
	"base-site-api/utils"
	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
)

// Handler for the uploads
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
	s := c.Params("type")

	categories, err := h.service.UploadCategories(s)

	if err != nil {
		log.Debugf("Error while getting upload categories by type slug %s", err)
		h.Error(c, 404)

		return
	}

	h.JSON(c, 200, categories)
}

func (h *Handler) ListUploads(c *fiber.Ctx) {
	s := c.Params("upload-category")
	page, size := utils.ParsePagination(c)

	uploads, count, err := h.service.UploadsByCategory(s, page, size)

	if err != nil {
		log.Debugf("Error while getting upload by category slug %s", err)
		h.Error(c, 404)

		return
	}

	p := h.CalculatePagination(page, size, count)

	h.JSON(c, 200, PaginatedUploads{
		p,
		uploads,
	})
}

func (h *Handler) Upload(c *fiber.Ctx) {
	file, err := c.FormFile("file")
	s := c.Params("upload-category")
	t := c.Params("type")

	if err != nil {
		log.Debugf("Error while parsing upload %s", err)
		h.Error(c, 400)

		return
	}

	r, err := h.service.Store(file, s, t)

	if err != nil {
		log.Debugf("Error while upload %s", err)
		h.Error(c, 400)

		return
	}

	h.JSON(c, 200, r)
}

func (h *Handler) CreateCategory(c *fiber.Ctx) {
	category := &models.UploadCategory{}
	t := c.Params("type")

	err := c.BodyParser(category)

	if err != nil {
		log.Debugf("Error while parsing upload category %s", err)
		h.Error(c, 400)

		return
	}

	id, err := h.service.StoreCategory(category.Name, category.SubPath, t)

	if err != nil {
		log.Errorf("Error while creating upload category: %s", err)
		h.Error(c, 500)

		return
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      id,
	}

	h.JSON(c, 201, &r)
}

func (h *Handler) UpdateCategory(c *fiber.Ctx) {
	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Error while parsing update upload category ID: %s", c.Params("id"))
		h.Error(c, 400)

		return
	}

	category := &models.UploadCategory{}

	err = c.BodyParser(category)

	if err != nil {
		log.Debugf("Error while parsing upload category %s", err)
		h.Error(c, 400)

		return
	}

	err = h.service.UpdateCategory(category.Name, category.SubPath, id)

	if err != nil {
		log.Errorf("Error while updating upload category: %s", err)
		h.Error(c, 500)

		return
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      id,
	}

	h.JSON(c, 201, &r)
}

func (h *Handler) Update(c *fiber.Ctx) {
	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Error while parsing update upload ID: %s", c.Params("id"))
		h.Error(c, 400)

		return
	}

	upload := &models.Upload{}

	err = c.BodyParser(upload)

	if err != nil {
		log.Debugf("Error while parsing upload %s", err)
		h.Error(c, 400)

		return
	}

	err = h.service.Update(upload.Description, id)

	if err != nil {
		log.Errorf("Error while updating upload: %s", err)

		h.Error(c, 500)

		return
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      id,
	}

	h.JSON(c, 201, &r)
}

// Remove handle deleting uploads
func (h *Handler) Remove(c *fiber.Ctx) {
	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Error while parsing upload id remove: %s", err)
		h.Error(c, 400)

		return
	}

	err = h.service.Delete(id)

	if err != nil {
		log.Errorf("Error while removing upload: %s", err)
		h.Error(c, 500)

		return
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      id,
	}

	h.JSON(c, 200, &r)
}

// Remove handle deleting uploads category
func (h *Handler) RemoveCategory(c *fiber.Ctx) {
	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Error while parsing upload category id remove: %s", err)
		h.Error(c, 400)

		return
	}

	err = h.service.DeleteCategory(id)

	if err != nil {
		log.Errorf("Error while removing upload category: %s", err)
		h.Error(c, 500)

		return
	}

	r := responses.SuccessResponse{
		Success: true,
		ID:      id,
	}

	h.JSON(c, 200, &r)
}
