package utils

import (
	"github.com/gofiber/fiber"
	"strconv"
)

// ParsePagination helper function which parse pagination from query params or return default
func ParsePagination(c *fiber.Ctx) (int, int) {
	p := c.Query("p")
	s := c.Query("s")
	var page int
	var size int

	if p == "" {
		page = 1
	} else {
		x, err := strconv.ParseInt(p, 10, 32)
		if err != nil {
			page = 1
		}

		page = int(x)
	}

	if s == "" {
		size = 10
	} else {
		x, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			size = 10
		}

		page = int(x)
	}

	return page, size
}
