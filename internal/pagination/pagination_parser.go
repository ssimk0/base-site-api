package pagination

import (
	"strconv"
)

// ParsePagination helper function which parse pagination from query params or return default
func ParsePagination(p string, s string) (int, int) {
	page := 1
	size := 10

	if p != "" {
		x, err := strconv.ParseInt(p, 10, 32)
		if err == nil {
			page = int(x)
		}
	}

	if s != "" {
		x, err := strconv.ParseInt(s, 10, 32)
		if err == nil {
			size = int(x)
		}

	}

	return page, size
}
