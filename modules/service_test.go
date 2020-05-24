package modules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_CalculateLimitAndOffset(t *testing.T) {
	s := Service{}

	offset, limit := s.CalculateLimitAndOffset(10, 10)

	assert.Equal(t, offset, 90)
	assert.Equal(t, limit, 10)
}

func TestService_CalculateLimitAndOffset_Page_1(t *testing.T) {
	s := Service{}

	offset, limit := s.CalculateLimitAndOffset(1, 10)

	assert.Equal(t, offset, 0)
	assert.Equal(t, limit, 10)
}
