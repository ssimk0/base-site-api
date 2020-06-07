package page

import (
	"base-site-api/models"
	"base-site-api/modules"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PageTestSuite struct {
	modules.RepositoryTestSuite
}

func (s *PageTestSuite) SetupTest() {
	s.Setup()
	s.Conn.Debug().AutoMigrate(
		&models.Page{},
		&models.PageCategory{},
	)
}

func (s *PageTestSuite) getTestPage() *models.Page {
	return &models.Page{
		Title:      "Test",
		Body:       "Body",
		Slug:       "test",
		UserID:     uint(1),
		CategoryID: uint(1),
	}
}

func (s *PageTestSuite) getTestPageCategory() *models.PageCategory {
	return &models.PageCategory{
		Name: "Test",
		Slug: "test",
	}
}

func (s *PageTestSuite) prepareTestData() ([]*models.Page, []*models.PageCategory) {
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

	categories, err := r.FindCategories()

	if err != nil {
		s.T().Errorf("Error List page categories %s", err)
	}

	assert.Len(s.T(), categories, len(categories))

	assert.Equal(s.T(), categories[0].Name, categories[0].Name)
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

	pages, err := r.FindAll(c[0].Slug)

	if err != nil {
		s.T().Errorf("Error find page category by slug %s", err)
	}

	assert.NotNil(s.T(), pages)
	assert.Len(s.T(), pages, len(p))
	assert.Equal(s.T(), p[0].Title, pages[0].Title)
}

func (s *PageTestSuite) TestFindBySlug() {
	p, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	page, err := r.FindBySlug(p[0].Slug)

	if err != nil {
		s.T().Errorf("Error find page category by slug %s", err)
	}

	assert.NotNil(s.T(), page)
	assert.Equal(s.T(), p[0].Title, page.Title)
}

func (s *PageTestSuite) TestUpdate() {
	p, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	data := &models.Page{
		Title: "new title",
	}

	err := r.Update(data, p[0].ID)

	if err != nil {
		s.T().Errorf("Error find page category by slug %s", err)
	}
	page := &models.Page{}
	s.Conn.First(page, p[0].ID)

	assert.NotNil(s.T(), page)
	assert.Equal(s.T(), data.Title, page.Title)
}

func (s *PageTestSuite) TestStore() {
	r := NewRepository(s.Conn)
	page := s.getTestPage()

	id, err := r.Store(page, uint(1))

	if err != nil {
		s.T().Errorf("Error find page category by slug %s", err)
	}
	p := &models.Page{}
	s.Conn.First(p, id)

	assert.NotNil(s.T(), p)
	assert.Equal(s.T(), page.Title, p.Title)
}

func (s *PageTestSuite) TestDelete() {
	p, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	err := r.Delete(p[0].ID)

	if err != nil {
		s.T().Errorf("Error Deleting article %s", err)
	}

	a := &models.Article{}
	s.Conn.First(a, p[0].ID)
	// Not found
	assert.Equal(s.T(), uint(0), a.ID)
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PageTestSuite))
}
