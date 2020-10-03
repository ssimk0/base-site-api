package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RandomStringTestSuite struct {
	suite.Suite
}

func (s *RandomStringTestSuite) TestGeneratingStringsLength() {

	assert.Len(s.T(), String(5), 5)
	assert.Len(s.T(), String(8), 8)
	assert.Len(s.T(), String(9), 9)
}

func (s *RandomStringTestSuite) TestGeneratingStringsDifference() {

	assert.NotEqual(s.T(), String(6), String(6))
	assert.NotEqual(s.T(), String(8), String(8))
}

func TestRandomStringTestSuite(t *testing.T) {
	suite.Run(t, new(RandomStringTestSuite))
}
