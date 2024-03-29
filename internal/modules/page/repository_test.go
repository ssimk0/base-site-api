package page

import (
	models2 "base-site-api/internal/models"
	"base-site-api/internal/tests/test_helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PageTestSuite struct {
	test_helper.RepositoryTestSuite
}

func (s *PageTestSuite) SetupTest() {
	s.Setup()
	s.Conn.Debug().AutoMigrate(
		&models2.Page{},
		&models2.PageCategory{},
	)
}

func (s *PageTestSuite) getTestPage() *models2.Page {
	return &models2.Page{
		Title:      "Test",
		Body:       "Body",
		Slug:       "test",
		UserID:     uint(1),
		CategoryID: uint(1),
	}
}

func (s *PageTestSuite) prepareTestData() ([]*models2.Page, []*models2.PageCategory) {
	pages := []*models2.Page{
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

	categories := []*models2.PageCategory{
		{
			Name: "Oznamy",
			Slug: "oznamy",
		},
		{
			Name: "Sluzby",
			Slug: "sluzby",
		},
	}

	for _, a := range categories {
		if err := s.Conn.Create(a).Error; err != nil {
			panic(err)
		}
	}

	for _, a := range pages {
		a.Category = *categories[0]
		if err := s.Conn.Create(a).Error; err != nil {
			panic(err)
		}
	}

	return pages, categories
}

func (s *PageTestSuite) TestFindCategories() {
	_, categories := s.prepareTestData()
	r := NewRepository(s.Conn)

	c, err := r.FindCategories()

	if err != nil {
		s.T().Errorf("Error List page categories %s", err)
	}

	assert.Len(s.T(), c, len(categories))

	assert.Equal(s.T(), c[0].Name, categories[0].Name)
}

func (s *PageTestSuite) TestCategoryBySlug() {
	_, c := s.prepareTestData()
	r := NewRepository(s.Conn)

	category, err := r.FindCategoryBySlug(c[0].Slug)

	if err != nil {
		s.T().Errorf("Error find page category by slug %s", err)
	}

	assert.NotNil(s.T(), category)
	assert.Equal(s.T(), c[0].Name, category.Name)
}

func (s *PageTestSuite) TestFindAll() {
	p, c := s.prepareTestData()
	r := NewRepository(s.Conn)

	pages, err := r.FindAllByCategorySlug(c[0].Slug)

	if err != nil {
		s.T().Errorf("Error find page category by slug %s", err)
	}

	assert.NotNil(s.T(), pages)
	assert.Len(s.T(), pages, len(p))
	assert.Equal(s.T(), p[0].Title, pages[0].Title)
}

func (s *PageTestSuite) TestFindAllNotFound() {
	r := NewRepository(s.Conn)

	_, err := r.FindAllByCategorySlug("not-found")

	assert.NotNil(s.T(), err)
}

func (s *PageTestSuite) TestFindBySlug() {
	p, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	page, _, err := r.FindBySlug(p[0].Slug)

	if err != nil {
		s.T().Errorf("Error find page category by slug %s", err)
	}

	assert.NotNil(s.T(), page)
	assert.Equal(s.T(), p[0].Title, page.Title)
}

func (s *PageTestSuite) TestFindBySlugNotFound() {
	r := NewRepository(s.Conn)

	_, _, err := r.FindBySlug("not-found")

	assert.NotNil(s.T(), err)

}

func (s *PageTestSuite) TestUpdate() {
	p, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	data := &models2.Page{
		Title: "new title",
		Body:  "other",
		Slug:  "new-title",
	}

	err := r.Update(data, p[0].ID)

	if err != nil {
		s.T().Errorf("Error find page category by slug %s", err)
	}
	page := &models2.Page{}
	s.Conn.First(page, p[0].ID)

	assert.NotNil(s.T(), page)
	assert.Equal(s.T(), data.Title, page.Title)
	assert.Equal(s.T(), data.Body, page.Body)
	assert.Equal(s.T(), data.Slug, page.Slug)
}

func (s *PageTestSuite) TestUpdateNotFound() {
	r := NewRepository(s.Conn)

	data := &models2.Page{
		Title: "new title",
		Body:  "other",
		Slug:  "new-title",
	}

	err := r.Update(data, 0)

	assert.NotNil(s.T(), err)
}

func (s *PageTestSuite) TestStore() {
	s.prepareTestData()
	r := NewRepository(s.Conn)
	page := s.getTestPage()

	id, err := r.Store(page, "sluzby", uint(1))

	if err != nil {
		s.T().Errorf("Error find page category by slug %s", err)
	}
	p := &models2.Page{}
	s.Conn.First(p, id)

	assert.NotNil(s.T(), p)
	assert.Equal(s.T(), page.Title, p.Title)
}

func (s *PageTestSuite) TestDelete() {
	p, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	err := r.Delete(p[0].ID, 1)

	if err != nil {
		s.T().Errorf("Error Deleting article %s", err)
	}

	a := &models2.Article{}
	s.Conn.First(a, p[0].ID)
	// Not found
	assert.Equal(s.T(), uint(0), a.ID)
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PageTestSuite))
}
