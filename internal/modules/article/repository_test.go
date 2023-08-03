package article

import (
	"base-site-api/internal/log"
	"base-site-api/internal/models"
	"base-site-api/internal/tests/test_helper"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ArticleTestSuite struct {
	test_helper.RepositoryTestSuite
}

func (s *ArticleTestSuite) SetupTest() {
	s.Setup()
	s.Conn.Debug().AutoMigrate(
		&models.Article{},
	)
}

func (s *ArticleTestSuite) getTestArticle() *models.Article {
	return &models.Article{
		Title:     "Test",
		Body:      "Body",
		Short:     "Short",
		Slug:      "test",
		Published: true,
	}
}

func (s *ArticleTestSuite) prepareTestData() []*models.Article {
	articles := []*models.Article{
		{
			Title:     "Test",
			Body:      "Body",
			Short:     "Short",
			Slug:      "test1",
			Published: true,
		},
		{
			Title:     "Test 2",
			Body:      "Body 2",
			Short:     "Short",
			Slug:      "test2",
			Published: true,
			Viewed:    2,
		},
		{
			Title:     "Test 3 ",
			Body:      "Body 3",
			Short:     "Short",
			Slug:      "test3",
			Published: true,
		},
	}

	for _, a := range articles {
		if err := s.Conn.Create(a).Error; err != nil {
			panic(err)
		}
	}

	return articles
}

func (s *ArticleTestSuite) TestStore() {
	a := s.getTestArticle()
	r := NewRepository(s.Conn)

	id, err := r.Store(a, 1)

	if err != nil {
		s.T().Errorf("Error store article %s", err)
	}
	log.Printf("ID store %d", id)
	assert.NotEqual(s.T(), 0, id)
}

func (s *ArticleTestSuite) TestFindAll() {
	data := s.prepareTestData()
	r := NewRepository(s.Conn)

	articles, count, err := r.FindAll("created_at", 0, 10)

	if err != nil {
		s.T().Errorf("Error List article %s", err)
	}

	assert.Len(s.T(), articles, len(data))

	assert.Equal(s.T(), articles[0].Title, data[2].Title)
	assert.Equal(s.T(), count, 3)
}

func (s *ArticleTestSuite) TestFindAllOrderViewed() {
	data := s.prepareTestData()
	r := NewRepository(s.Conn)

	articles, count, err := r.FindAll("viewed", 1, 10)

	if err != nil {
		s.T().Errorf("Error List article %s", err)
	}

	assert.Len(s.T(), articles, len(data))

	assert.Equal(s.T(), articles[0].Slug, "test2")
	assert.Equal(s.T(), count, 3)
}

func (s *ArticleTestSuite) TestFind() {
	data := s.prepareTestData()
	r := NewRepository(s.Conn)

	article, err := r.Find(data[0].ID)

	if err != nil {
		s.T().Errorf("Error Find article %s", err)
	}

	assert.Equal(s.T(), article.Title, data[0].Title)
}

func (s *ArticleTestSuite) TestFindBySlug() {
	data := s.prepareTestData()
	r := NewRepository(s.Conn)

	article, err := r.FindBySlug(data[0].Slug)

	if err != nil {
		s.T().Errorf("Error FindBySlug article %s", err)
	}

	assert.Equal(s.T(), article.Title, data[0].Title)
}

func (s *ArticleTestSuite) TestUpdate() {
	data := s.prepareTestData()
	r := NewRepository(s.Conn)

	article := data[0]
	article.Title = "New Title"

	err := r.Update(article, article.ID)

	if err != nil {
		s.T().Errorf("Error Update article %s", err)
	}
	a := &models.Article{}
	s.Conn.First(a, article.ID)

	assert.Equal(s.T(), a.Title, "New Title")
}

func (s *ArticleTestSuite) TestDelete() {
	data := s.prepareTestData()
	r := NewRepository(s.Conn)

	err := r.Delete(data[0].ID, 1)

	if err != nil {
		s.T().Errorf("Error Deleting article %s", err)
	}

	a := &models.Article{}
	s.Conn.First(a, data[0].ID)
	// Not found
	assert.Equal(s.T(), uint(0), a.ID)
}

func TestArticleRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ArticleTestSuite))
}
