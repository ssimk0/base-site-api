package handlers

import (
	"base-site-api/internal/app/dto"
	"base-site-api/internal/app/models"
	"base-site-api/internal/log"
	"base-site-api/internal/modules/upload"
	"base-site-api/internal/pagination"
	"github.com/gofiber/fiber/v2"
)

// PageHandler for the upload
type UploadHandler struct {
	Handler
	service upload.Service
}

func NewUploadHandler(s upload.Service) *UploadHandler {
	return &UploadHandler{
		service: s,
	}
}

func (h *UploadHandler) ListCategories(c *fiber.Ctx) error {
	s := c.Params("type")

	categories, err := h.service.UploadCategories(s)

	if err != nil {
		log.Debugf("Error while getting upload categories by type slug %s", err)
		return h.Error(404)
	}

	return h.JSON(c, 200, categories)
}

func (h *UploadHandler) ListUploads(c *fiber.Ctx) error {
	s := c.Params("uploadCategory")
	page, size := pagination.ParsePagination(c)

	uploads, count, err := h.service.UploadsByCategory(s, page, size)

	if err != nil {
		log.Debugf("Error while getting upload by category slug %s", err)
		return h.Error(404)
	}

	p := h.CalculatePagination(page, size, count)

	return h.JSON(c, 200, upload.PaginatedUploads{
		Pagination: p,
		Uploads:    uploads,
	})
}

func (h *UploadHandler) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	s := c.Params("uploadCategory")
	t := c.Params("type")

	if err != nil {
		log.Debugf("Error while parsing upload %s", err)
		return h.Error(400)
	}

	r, err := h.service.Store(file, s, t)

	if err != nil {
		log.Debugf("Error while upload %s", err)
		return h.Error(400)
	}

	return h.JSON(c, 200, r)
}

func (h *UploadHandler) CreateCategory(c *fiber.Ctx) error {
	category := &models.UploadCategory{}
	t := c.Params("type")

	err := c.BodyParser(category)

	if err != nil {
		log.Debugf("Error while parsing upload category %s", err)
		return h.Error(400)
	}

	id, err := h.service.StoreCategory(category.Name, category.SubPath, category.Description, t)

	if err != nil {
		log.Errorf("Error while creating upload category: %s", err)
		return h.Error(500)
	}

	r := dto.SuccessResponse{
		Success: true,
		ID:      id,
	}

	return h.JSON(c, 201, &r)
}

func (h *UploadHandler) UpdateCategory(c *fiber.Ctx) error {
	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Error while parsing update upload category ID: %s", c.Params("id"))
		return h.Error(400)
	}

	category := &models.UploadCategory{}

	err = c.BodyParser(category)

	if err != nil {
		log.Debugf("Error while parsing upload category %s", err)
		return h.Error(400)
	}

	err = h.service.UpdateCategory(category.Name, category.SubPath, id)

	if err != nil {
		log.Errorf("Error while updating upload category: %s", err)
		return h.Error(500)
	}

	r := dto.SuccessResponse{
		Success: true,
		ID:      id,
	}

	return h.JSON(c, 201, &r)
}

func (h *UploadHandler) Update(c *fiber.Ctx) error {
	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Error while parsing update upload ID: %s", c.Params("id"))
		return h.Error(400)
	}

	u := &models.Upload{}

	err = c.BodyParser(u)

	if err != nil {
		log.Debugf("Error while parsing upload %s", err)
		return h.Error(400)
	}

	err = h.service.Update(u.Description, id)

	if err != nil {
		log.Errorf("Error while updating upload: %s", err)

		return h.Error(500)
	}

	r := dto.SuccessResponse{
		Success: true,
		ID:      id,
	}

	return h.JSON(c, 201, &r)
}

// Remove handle deleting upload
func (h *UploadHandler) Remove(c *fiber.Ctx) error {
	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Error while parsing upload id remove: %s", err)
		return h.Error(400)
	}

	err = h.service.Delete(id)

	if err != nil {
		log.Errorf("Error while removing upload: %s", err)
		return h.Error(500)
	}

	r := dto.SuccessResponse{
		Success: true,
		ID:      id,
	}

	return h.JSON(c, 200, &r)
}

// Remove handle deleting upload category
func (h *UploadHandler) RemoveCategory(c *fiber.Ctx) error {
	id, err := h.ParseID(c)

	if err != nil {
		log.Debugf("Error while parsing upload category id remove: %s", err)
		return h.Error(400)
	}

	err = h.service.DeleteCategory(id)

	if err != nil {
		log.Errorf("Error while removing upload category: %s", err)
		return h.Error(500)
	}

	r := dto.SuccessResponse{
		Success: true,
		ID:      id,
	}

	return h.JSON(c, 200, &r)
}
