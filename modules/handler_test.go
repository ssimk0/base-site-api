package modules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_CalculatePagination(t *testing.T) {
	h := Handler{}

	p := h.CalculatePagination(1, 10, 100)

	assert.Equal(t, p.Page, 1)
	assert.Equal(t, p.TotalPages, float64(10))
	assert.Equal(t, p.Total, 100)
	assert.Equal(t, p.PageSize, 10)
}

func TestHandler_CalculatePagination_1_Page(t *testing.T) {
	h := Handler{}

	p := h.CalculatePagination(1, 10, 5)

	assert.Equal(t, p.Page, 1)
	assert.Equal(t, p.TotalPages, float64(1))
	assert.Equal(t, p.Total, 5)
	assert.Equal(t, p.PageSize, 10)
}
