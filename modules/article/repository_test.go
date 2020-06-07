package article

import (
	"base-site-api/models"
	"base-site-api/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ArticleTestSuite struct {
	suite.Suite
	conn        *gorm.DB
	cleanupHook func()
}

func (s *ArticleTestSuite) SetupTest() {
	var err error
	s.conn, err = gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		log.Fatal(err)
	}

	s.conn.LogMode(true)
	s.conn.Debug().AutoMigrate(
		&models.Article{},
	)
}

func (s *ArticleTestSuite) BeforeTest(suiteName, testName string) {
	log.Debugf("Before test {} from suite {}", suiteName, testName)
	s.cleanupHook = utils.DeleteCreatedEntities(s.conn)
}

func (s *ArticleTestSuite) AfterTest(suiteName, testName string) {
	log.Debugf("After test {} from suite {}", suiteName, testName)
	s.cleanupHook()
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
		if err := s.conn.Create(a).Error; err != nil {
			panic(err)
		}
	}

	return articles
}

func (s *ArticleTestSuite) TestStore() {
	a := s.getTestArticle()
	r := NewRepository(s.conn)

	id, err := r.Store(a, 1)

	if err != nil {
		s.T().Errorf("Error store article %s", err)
	}
	log.Printf("ID store %d", id)
	assert.NotEqual(s.T(), 0, id)
}

func (s *ArticleTestSuite) TestFindAll() {
	data := s.prepareTestData()
	r := NewRepository(s.conn)

	articles, count, err := r.FindAll("created_at", 0, 10)

	if err != nil {
		s.T().Errorf("Error List article %s", err)
	}

	assert.Len(s.T(), articles, len(data))

	assert.Equal(s.T(), articles[0].Title, data[0].Title)
	assert.Equal(s.T(), count, 3)
}

func (s *ArticleTestSuite) TestFindAllOrderViewed() {
	data := s.prepareTestData()
	r := NewRepository(s.conn)

	articles, count, err := r.FindAll("viewed desc", 0, 10)

	if err != nil {
		s.T().Errorf("Error List article %s", err)
	}

	assert.Len(s.T(), articles, len(data))

	assert.Equal(s.T(), articles[0].Slug, "test2")
	assert.Equal(s.T(), count, 3)
}

func (s *ArticleTestSuite) TestFind() {
	data := s.prepareTestData()
	r := NewRepository(s.conn)

	article, err := r.Find(data[0].ID)

	if err != nil {
		s.T().Errorf("Error Find article %s", err)
	}

	assert.Equal(s.T(), article.Title, data[0].Title)
}

func (s *ArticleTestSuite) TestFindBySlug() {
	data := s.prepareTestData()
	r := NewRepository(s.conn)

	article, err := r.FindBySlug(data[0].Slug)

	if err != nil {
		s.T().Errorf("Error FindBySlug article %s", err)
	}

	assert.Equal(s.T(), article.Title, data[0].Title)
}

func (s *ArticleTestSuite) TestUpdate() {
	data := s.prepareTestData()
	r := NewRepository(s.conn)

	article := data[0]
	article.Title = "New Title"

	err := r.Update(article, article.ID)

	if err != nil {
		s.T().Errorf("Error Update article %s", err)
	}
	a := &models.Article{}
	s.conn.First(a, article.ID)

	assert.Equal(s.T(), a.Title, "New Title")
}

func (s *ArticleTestSuite) TestDelete() {
	data := s.prepareTestData()
	r := NewRepository(s.conn)

	err := r.Delete(data[0].ID)

	if err != nil {
		s.T().Errorf("Error Deleting article %s", err)
	}

	a := &models.Article{}
	s.conn.First(a, data[0].ID)
	// Not found
	assert.Equal(s.T(), uint(0), a.ID)
}

func TestArticleRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ArticleTestSuite))
}
