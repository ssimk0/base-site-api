package utils

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RandomStringTestSuite struct {
	suite.Suite
}

func (s *RandomStringTestSuite) TestGeneratingStringsLength() {

	assert.Len(s.T(), GenerateRandomString(5), 5)
	assert.Len(s.T(), GenerateRandomString(8), 8)
	assert.Len(s.T(), GenerateRandomString(9), 9)
}

func (s *RandomStringTestSuite) TestGeneratingStringsDifference() {

	assert.NotEqual(s.T(), GenerateRandomString(6), GenerateRandomString(6))
	assert.NotEqual(s.T(), GenerateRandomString(8), GenerateRandomString(8))
}

func TestRandomStringTestSuite(t *testing.T) {
	suite.Run(t, new(RandomStringTestSuite))
}
