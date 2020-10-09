package pagination

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaginationParsing_default(t *testing.T) {
	page, size := ParsePagination("", "")

	assert.Equal(t, page, 1)
	assert.Equal(t, size, 10)
}

func TestPaginationParsing_with_page(t *testing.T) {
	page, size := ParsePagination("5", "")

	assert.Equal(t, page, 5)
	assert.Equal(t, size, 10)
}

func TestPaginationParsing_with_size(t *testing.T) {
	page, size := ParsePagination("", "20")

	assert.Equal(t, page, 1)
	assert.Equal(t, size, 20)
}

func TestPaginationParsing_with_both(t *testing.T) {
	page, size := ParsePagination("5", "15")

	assert.Equal(t, page, 5)
	assert.Equal(t, size, 15)
}

func TestPaginationParsing_invalid_page(t *testing.T) {
	page, size := ParsePagination("x", "")

	assert.Equal(t, page, 1)
	assert.Equal(t, size, 10)
}

func TestPaginationParsing_invalid_size(t *testing.T) {
	page, size := ParsePagination("", "x")

	assert.Equal(t, page, 1)
	assert.Equal(t, size, 10)
}
