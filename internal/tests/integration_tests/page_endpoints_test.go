package integration_tests

import (
	"base-site-api/internal/app/models"
	"base-site-api/internal/modules/page"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PageEndpointsTestSuite struct {
	EndpointsTestSuite
}

func (s *PageEndpointsTestSuite) SetupTest() {
	s.SetupApp()
}

func (s *PageEndpointsTestSuite) prepareTestData() ([]*models.Page, []*models.PageCategory) {
	pages := []*models.Page{
		{
			Title: "Test",
			Body:  "Body",
			Slug:  "test1",
		},
		{
			Title: "Test 2",
			Body:  "Body 2",
			Slug:  "test2",
		},
		{
			Title: "Test 3 ",
			Body:  "Body 3",
			Slug:  "test3",
		},
	}

	categories := []*models.PageCategory{
		{
			Name: "Oznamy",
			Slug: "oznamy",
		},
		{
			Name: "Sluzby",
			Slug: "sluzby",
		},
	}

	for _, c := range categories {
		if err := s.Conn.Create(c).Error; err != nil {
			panic(err)
		}
	}

	for _, p := range pages {
		p.Category = *categories[0]
		if err := s.Conn.Create(p).Error; err != nil {
			panic(err)
		}
	}

	return pages, categories
}

func (s *PageEndpointsTestSuite) TestFindPageCategories() {

	_, c := s.prepareTestData()

	req, _ := http.NewRequest(
		"GET",
		"/api/v1/pages/",
		nil,
	)

	// Perform the request plain with the app.
	// The -1 disables request latency.
	res, err := s.app.Test(req, -1)

	if err != nil {
		s.T().Errorf("Error while calling /api/v1/pages %s", err)
	}

	// // verify that no error occurred, that is not expected
	assert.Equal(s.T(), res.StatusCode, 200)

	d := json.NewDecoder(res.Body)
	var resData []*models.PageCategory
	err = d.Decode(&resData)

	if err != nil {
		s.T().Errorf("Error while decoding response of /api/v1/pages %s", err)
	}

	assert.Equal(s.T(), resData[0].Slug, c[0].Slug)
	assert.Equal(s.T(), resData[1].Slug, c[1].Slug)
}

func (s *PageEndpointsTestSuite) TestFindPagesOfCategory() {

	p, c := s.prepareTestData()

	req, _ := http.NewRequest(
		"GET",
		"/api/v1/pages/"+c[0].Slug,
		nil,
	)

	// Perform the request plain with the app.
	// The -1 disables request latency.
	res, err := s.app.Test(req, -1)

	if err != nil {
		s.T().Errorf("Error while calling /api/v1/pages/oznamy %s", err)
	}

	// // verify that no error occurred, that is not expected
	assert.Equal(s.T(), res.StatusCode, 200)

	d := json.NewDecoder(res.Body)
	var resData []*page.PageDetail
	err = d.Decode(&resData)

	if err != nil {
		s.T().Errorf("Error while decoding response of /api/v1/pages/oznamy %s", err)
	}

	assert.Len(s.T(), resData, 3)
	assert.Equal(s.T(), resData[0].Slug, p[0].Slug)
	assert.Equal(s.T(), resData[1].Slug, p[1].Slug)
}

func (s *PageEndpointsTestSuite) TestFindEmptyPagesOfCategory() {

	s.prepareTestData()

	req, _ := http.NewRequest(
		"GET",
		"/api/v1/pages/sluzby",
		nil,
	)

	// Perform the request plain with the app.
	// The -1 disables request latency.
	res, err := s.app.Test(req, -1)

	if err != nil {
		s.T().Errorf("Error while calling /api/v1/pages/sluzby %s", err)
	}

	// // verify that no error occurred, that is not expected
	assert.Equal(s.T(), res.StatusCode, 200)

	d := json.NewDecoder(res.Body)
	var resData []*models.Page
	err = d.Decode(&resData)

	if err != nil {
		s.T().Errorf("Error while decoding response of /api/v1/pages/sluzby %s", err)
	}

	assert.Len(s.T(), resData, 0)
}

func (s *PageEndpointsTestSuite) TestFindNotFoundPagesOfCategory() {

	s.prepareTestData()

	req, _ := http.NewRequest(
		"GET",
		"/api/v1/pages/other",
		nil,
	)

	// Perform the request plain with the app.
	// The -1 disables request latency.
	res, err := s.app.Test(req, -1)

	if err != nil {
		s.T().Errorf("Error while calling /api/v1/pages/other %s", err)
	}

	// // verify that no error occurred, that is not expected
	assert.Equal(s.T(), res.StatusCode, 404)
}

func TestEndpointsTestSuite(t *testing.T) {
	suite.Run(t, new(PageEndpointsTestSuite))
}
