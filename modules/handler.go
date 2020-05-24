package modules

import "math"

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
