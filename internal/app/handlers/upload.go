package handlers

import (
	"base-site-api/internal/app/dto"
	"base-site-api/internal/log"
	"base-site-api/internal/models"
	"base-site-api/internal/modules/upload"
	"base-site-api/internal/pagination"
	"base-site-api/internal/random"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
)

// PageHandler for the upload
type UploadHandler struct {
	Handler
	service    upload.Service
	repository upload.Repository
}

func NewUploadHandler(s upload.Service) *UploadHandler {
	return &UploadHandler{
		service: s,
	}
}

func (h *UploadHandler) ListCategories(c *fiber.Ctx) error {
	s := c.Params("type")

	categories, err := h.repository.FindCategoriesByType(s)

	if err != nil {
		log.Debugf("Error while getting upload categories by type slug %s", err)
		return h.Error(404)
	}

	return h.JSON(c, 200, categories)
}

func (h *UploadHandler) DownloadUpload(c *fiber.Ctx) error {
	id := c.Params("id")

	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return h.Error(400)
	}

	u, err := h.repository.Find(uint(uid))

	if err != nil {
		log.Debugf("Error while getting latest upload by type slug %s", err)
		return h.Error(404)
	}
	file := path.Join("/tmp", random.String(6))
	err = downloadFile(file, u.File)

	if err != nil {
		return h.Error(500)
	}

	return c.SendFile(file, false)
}

func (h *UploadHandler) Detail(c *fiber.Ctx) error {
	id := c.Params("id")

	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return h.Error(400)
	}

	u, err := h.repository.Find(uint(uid))

	if err != nil {
		log.Debugf("Error while getting latest upload by type slug %s", err)
		return h.Error(404)
	}

	return c.JSON(u)
}

func (h *UploadHandler) EditUpload(c *fiber.Ctx) error {
	id := c.Params("id")

	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return h.Error(400)
	}

	u := &models.Upload{}

	err = c.BodyParser(u)

	if err != nil {
		return h.Error(400)
	}

	err = h.repository.Update(u.Description, uint(uid))

	if err != nil {
		log.Debugf("Error while edit upload %s", err)
		return h.Error(404)
	}

	return c.JSON(u)
}

func (h *UploadHandler) LastestUpload(c *fiber.Ctx) error {
	s := c.Params("uploadCategory")

	u, err := h.repository.FindLatestUploadByCategory(s)

	if err != nil {
		log.Debugf("Error while getting latest upload by type slug %s", err)
		return h.Error(404)
	}

	file := path.Join("/tmp", random.String(6))
	err = downloadFile(file, u.File)

	if err != nil {
		return h.Error(500)
	}

	return c.SendFile(file, false)
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func (h *UploadHandler) ListUploads(c *fiber.Ctx) error {
	s := c.Params("uploadCategory")
	qp := c.Query("p")
	qs := c.Query("s")
	page, size := pagination.ParsePagination(qp, qs)

	uploads, count, err := h.repository.FindUploadsByCategory(s, page, size)

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
	desc := c.FormValue("description", "")
	s := c.Params("uploadCategory")
	t := c.Params("type")

	if err != nil {
		log.Debugf("Error while parsing upload %s", err)
		return h.Error(400)
	}

	r, err := h.service.Store(file, desc, s, t)

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

	id, err := h.repository.StoreCategory(category.Name, category.SubPath, category.Description, t)

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

	err = h.repository.UpdateCategory(category.Name, category.SubPath, "", id)

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

	err = h.repository.Update(u.Description, id)

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

	err = h.repository.Delete(id)

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

	err = h.repository.DeleteCategory(id)

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
