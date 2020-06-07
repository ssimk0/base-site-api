package modules

import (
	"github.com/gofiber/fiber"
	"math"
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
