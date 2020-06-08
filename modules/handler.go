package modules

import (
	"github.com/gofiber/fiber"
	"math"
	"strconv"
)

type Handler struct {
}

func (h *Handler) CalculatePagination(page int, size int, count int) *Pagination {
	return &Pagination{
		Page:       page,
		PageSize:   size,
		Total:      count,
		TotalPages: math.Ceil(float64(count) / float64(size)),
	}
}

func (h *Handler) JSON(c *fiber.Ctx, status int, data interface{}) {
	if err := c.Status(status).JSON(data); err != nil {
		c.Next(err)
	}
}

func (h *Handler) Error(c *fiber.Ctx, status int) {
	c.Next(fiber.NewError(status))
}

func (h *Handler) ErrorWithMessage(c *fiber.Ctx, status int, message string) {
	c.Next(fiber.NewError(status, message))
}

func (h *Handler) ParseUserId(c *fiber.Ctx) uint {
	return c.Locals("userID").(uint)
}

func (h *Handler) ParseID(c *fiber.Ctx) (uint, error) {
	id := c.Params("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return uint(0), err
	}

	return uint(uid), err
}
